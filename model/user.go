package model

import "go.mongodb.org/mongo-driver/mongo"

type Users interface {
	get(name string, password string) (User, error)
	all() ([]User, error)
	create(theUser User) error
	update(theUser User) (User, error)
	delete(theUser User) (User, error)
	deleteAll() error
}

type userDB struct {
	collection *mongo.Collection
}

func CreateUserRepo(coll *mongo.Collection) Users {
	return &userDB{
		collection: coll,
	}
}

func (u *userDB) get(name string, password string) (User, error) {}
func (u *userDB) all() ([]User, error)                           {}
func (u *userDB) create(theUser User) error                      {}
func (u *userDB) update(theUser User) (User, error)              {}
func (u *userDB) delete(theUser User) (User, error)              {}
func (u *userDB) deleteAll() error                               {}
