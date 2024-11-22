package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tasks interface {
	Create(theTask Task) error
	Delete(theId string, theTask string, theDate string) (Task, error)
	DeleteAll(userid string) error
	Update(newTask Task, name string) (int, error)
	Get(theId string, theTask string, theDate string) (Task, error)
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

// Creating a new the tasks of a user
func (t *taskDB) Create(theTask Task) error {
	theTask.status = "pending"
	_, err := t.collection.InsertOne(context.TODO(), theTask)
	if err != nil {
		log.Fatal("error while creating a task")
		return err
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
	filter := bson.D{{Key: "user_id", Value: theDate}, {Key: "Name", Value: theTask}, {Key: "Date", Value: theDate}}
	var res Task
	err := t.collection.FindOne(context.Background(), filter).Decode(&res)
	if err != nil {
		return task, err
	}
	log.Println("deleted the task")
	return res, nil
}

// Fetching all the tasks of a user which are pending from the database
func (t *taskDB) All(theId string) ([]Task, error) {
	cursor, err := t.collection.Find(context.TODO(), bson.D{{Key: "user_id", Value: theId}, {Key: "status", Value: "pending"}})
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
	cursor, err := t.collection.Find(context.TODO(), bson.D{{Key: "user_id", Value: theId}, {Key: "status"}})
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
	filter := bson.D{{Key: "user_id", Value: theDate}, {Key: "Name", Value: theTask}, {Key: "Date", Value: theDate}}
	_, err := t.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return task, err
	}
	log.Println("deleted the task")
	return task, nil
}

// Deleting all the task of a user from the database
func (t *taskDB) DeleteAll(userid string) error {
	res, err := t.collection.DeleteMany(context.TODO(), bson.D{{Key: "user_id", Value: userid}})
	if err != nil {
		log.Fatal("error while creating a task")
		return err
	}
	log.Println("deleted all the tasks of the user", res.DeletedCount)
	return nil
}

// Updating the task of a user
func (t *taskDB) Update(newTask Task, name string) (int, error) {
	newTask.status = "pending"
	filter := bson.D{{Key: "user_id", Value: newTask.UserId}, {Key: "name", Value: name}, {Key: "Date", Value: newTask.Date}}
	update := bson.D{{Key: "$set", Value: newTask}}
	upd, err := t.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}
	log.Println("updated the user", upd.ModifiedCount)
	return int(upd.ModifiedCount), nil
}

// Marking the task of a user as done from the database
func (t *taskDB) Done(theId string, theTask string, theDate string) (Task, error) {
	completedTask := Task{
		UserId: theId,
		Name:   theTask,
		Date:   theDate,
		status: "Done",
	}
	filter := bson.D{{Key: "user_id", Value: theId}, {Key: "name", Value: theTask}, {Key: "Date", Value: theDate}}
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
		status: "Pending",
	}
	filter := bson.D{{Key: "user_id", Value: theId}, {Key: "name", Value: theTask}, {Key: "Date", Value: theDate}}
	update := bson.D{{Key: "$set", Value: pendingTask}}
	upd, err := t.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return pendingTask, err
	}
	log.Println("updated the user", upd.ModifiedCount)
	return pendingTask, nil
}
