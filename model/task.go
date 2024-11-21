package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Tasks interface {
	Create(theTask Task) error
	Delete(theId string, theTask string, theDate string) (Task, error)
	DeleteAll(userid string) error
	Update(newTask Task, name string) (Task, error)
	Get(theId string, theTask string, theDate string) (Task, error)
	Done(theId string, theTask string, theDate string) (Task, error)
	Undone(theId string, theTask string, theDate string) (Task, error)
	All(theId string) ([]Task, error)
}

type taskDB struct {
	collection *mongo.Collection
}

func CreateTaskRepo(coll *mongo.Collection) Tasks {
	return &taskDB{
		collection: coll,
	}
}

func (t *taskDB) Create(theTask Task) error                                         {}
func (t *taskDB) Delete(theId string, theTask string, theDate string) (Task, error) {}
func (t *taskDB) DeleteAll(userid string) error                                     {}
func (t *taskDB) Update(newTask Task, name string) (Task, error)                    {}
func (t *taskDB) Get(theId string, theTask string, theDate string) (Task, error)    {}
func (t *taskDB) Done(theId string, theTask string, theDate string) (Task, error)   {}
func (t *taskDB) Undone(theId string, theTask string, theDate string) (Task, error) {}
func (t *taskDB) All(theId string) ([]Task, error)                                  {}
