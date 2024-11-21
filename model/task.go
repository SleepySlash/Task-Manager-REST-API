package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Tasks interface {
	create(theTask Task) error
	delete(theTask string) (Task, error)
	deleteAll() error
	update(theTask string) (Task, error)
	get(theTask string) (Task, error)
	all() ([]Task, error)
}

type taskDB struct {
	collection *mongo.Collection
}

func CreateTaskRepo(coll *mongo.Collection) Tasks {
	return &taskDB{
		collection: coll,
	}
}

func (t *taskDB) create(theTask Task) error {

}

func (t *taskDB) delete(theTask string) (Task, error) {

}
func (t *taskDB) deleteAll() error {

}

func (t *taskDB) update(theTask string) (Task, error) {

}

func (t *taskDB) get(theTask string) (Task, error) {

}
func (t *taskDB) all() ([]Task, error) {

}
