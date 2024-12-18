package model

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Tasks interface {
	Create(theTask Task) error
	CreateMany(theTask []Task) error
	Delete(theId string, theTask string, theDate string) (Task, error)
	DeleteAll(userid string) (int64, error)
	Update(newTask Task, name string) (int, error)
	Get(theId string, theTask string, theDate string) (Task, error)
	AllDone(theId string, theTask []string) ([]Task, error)
	Done(theId string, theTask string, theDate string) (Task, error)
	Undone(theId string, theTask string, theDate string) (Task, error)
	All(theId string) ([]Task, error)
	AllTasks(theId string) ([]Task, error)
}

type taskDB struct {
	collection *mongo.Collection
}

func CreateTaskRepo(coll *mongo.Collection) Tasks {
	return &taskDB{
		collection: coll,
	}
}

// Creating a new task of a user
func (t *taskDB) Create(theTask Task) error {
	_, err := t.collection.InsertOne(context.TODO(), theTask)
	if err != nil {
		log.Fatal("error while creating a task")
		return err
	}
	return nil
}

// Creating a multiple new tasks of a user
func (t *taskDB) CreateMany(theTasks []Task) error {
	modelChannel := make(chan mongo.WriteModel, len(theTasks))

	var wg sync.WaitGroup
	worker := func(tasks <-chan Task) {
		defer wg.Done()
		for task := range tasks {
			modelChannel <- mongo.NewInsertOneModel().SetDocument(task)
		}
	}
	tasksChannel := make(chan Task, len(theTasks))
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(tasksChannel)
	}
	for _, task := range theTasks {
		tasksChannel <- task
	}
	close(tasksChannel)
	wg.Wait()
	close(modelChannel)
	var models []mongo.WriteModel
	for i := range modelChannel {
		models = append(models, i)
	}
	if len(models) > 0 {
		bulkWriteOptions := options.BulkWrite().SetOrdered(false)
		_, err := t.collection.BulkWrite(context.TODO(), models, bulkWriteOptions)
		if err != nil {
			log.Println("Error is storing the items into the DB (bulk write)")
			return err
		}
		log.Println("bulk write executed")
	}

	return nil
}

// Fetching a specific task of a user from the database
func (t *taskDB) Get(theId string, theTask string, theDate string) (Task, error) {
	task := Task{
		UserId: theId,
		Name:   theTask,
		Date:   theDate,
	}
	filter := bson.D{{Key: "userid", Value: theId}, {Key: "name", Value: theTask}, {Key: "date", Value: theDate}}
	var res Task
	err := t.collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return task, err
	}
	log.Println("found the task")
	return res, nil
}

// Fetching all the tasks of a user which are pending from the database
func (t *taskDB) All(theId string) ([]Task, error) {
	cursor, err := t.collection.Find(context.TODO(), bson.D{{Key: "userid", Value: theId}, {Key: "status", Value: "pending"}})
	if err != nil {
		log.Fatal("error while creating a task")
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var tasks []Task
	err = cursor.All(context.TODO(), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Fetching all the tasks of a user from the database
func (t *taskDB) AllTasks(theId string) ([]Task, error) {
	cursor, err := t.collection.Find(context.TODO(), bson.D{{Key: "userid", Value: theId}})
	if err != nil {
		log.Fatal("error while creating a task")
		return nil, err
	}
	defer cursor.Close(context.TODO())
	var tasks []Task
	err = cursor.All(context.TODO(), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// Deleting a specific task of a user from the database
func (t *taskDB) Delete(theId string, theTask string, theDate string) (Task, error) {
	task := Task{
		UserId: theId,
		Name:   theTask,
		Date:   theDate,
	}
	filter := bson.D{{Key: "userid", Value: theId}, {Key: "name", Value: theTask}, {Key: "date", Value: theDate}}
	_, err := t.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return task, err
	}
	log.Println("deleted the task")
	return task, nil
}

// Deleting all the task of a user from the database
func (t *taskDB) DeleteAll(userid string) (int64, error) {
	res, err := t.collection.DeleteMany(context.TODO(), bson.D{{Key: "userid", Value: userid}})
	if err != nil {
		log.Fatal("error while creating a task")
		return 0, err
	}
	log.Println("deleted all the tasks of the user", res.DeletedCount)
	return res.DeletedCount, nil
}

// Updating the task of a user
func (t *taskDB) Update(newTask Task, name string) (int, error) {
	filter := bson.D{{Key: "userid", Value: newTask.UserId}, {Key: "name", Value: name}, {Key: "date", Value: newTask.Date}}
	update := bson.D{{Key: "$set", Value: newTask}}
	updOption := options.Update().SetUpsert(true)
	upd, err := t.collection.UpdateOne(context.TODO(), filter, update, updOption)
	if err != nil {
		return 0, err
	}
	log.Println("updated the user", upd.ModifiedCount)
	return int(upd.ModifiedCount), nil
}

// Marking all the given tasks of a user as done from the database
func (t *taskDB) AllDone(theId string, theTasks []string) ([]Task, error) {

	modelsChannel := make(chan *mongo.UpdateOneModel, len(theTasks))
	var wg sync.WaitGroup
	worker := func(tasks <-chan string) {
		defer wg.Done()
		for task := range tasks {
			filter := bson.D{{Key: "userid", Value: theId}, {Key: "name", Value: task}}
			update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "Done"}}}}
			modelsChannel <- mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update)
		}
	}

	tasksChannel := make(chan string, len(theTasks))
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(tasksChannel)
	}

	for _, task := range theTasks {
		tasksChannel <- task
	}

	close(tasksChannel)
	wg.Wait()
	close(modelsChannel)

	var models []mongo.WriteModel
	for update := range modelsChannel {
		models = append(models, update)
	}
	bulkOptions := options.BulkWrite().SetOrdered(false)
	_, err := t.collection.BulkWrite(context.TODO(), models, bulkOptions)
	if err != nil {
		log.Println("error while bulk writing the tasks as done")
		return nil, err
	}
	filter := bson.D{
		{Key: "userid", Value: theId},
		{Key: "name", Value: bson.D{{Key: "$in", Value: theTasks}}},
		{Key: "status", Value: "Done"},
	}
	cursor, err := t.collection.Find(context.Background(), filter)
	if err != nil {
		log.Println("error while retrieving updated tasks:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())
	var updatedTasks []Task
	if err = cursor.All(context.Background(), &updatedTasks); err != nil {
		log.Println("error while decoding updated tasks:", err)
		return nil, err
	}

	return updatedTasks, nil

}

// Marking the task of a user as done from the database
func (t *taskDB) Done(theId string, theTask string, theDate string) (Task, error) {
	completedTask := Task{
		UserId: theId,
		Name:   theTask,
		Date:   theDate,
		Status: "Done",
	}
	filter := bson.D{{Key: "userid", Value: theId}, {Key: "name", Value: theTask}, {Key: "date", Value: theDate}}
	update := bson.D{{Key: "$set", Value: completedTask}}
	upd, err := t.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return completedTask, err
	}
	log.Println("updated the user", upd.ModifiedCount)
	return completedTask, nil

}

// Marking the task of a user as pending from the database
func (t *taskDB) Undone(theId string, theTask string, theDate string) (Task, error) {
	pendingTask := Task{
		UserId: theId,
		Name:   theTask,
		Date:   theDate,
		Status: "pending",
	}
	filter := bson.D{{Key: "userid", Value: theId}, {Key: "name", Value: theTask}, {Key: "date", Value: theDate}}
	update := bson.D{{Key: "$set", Value: pendingTask}}
	upd, err := t.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return pendingTask, err
	}
	log.Println("updated the user", upd.ModifiedCount)
	return pendingTask, nil
}
