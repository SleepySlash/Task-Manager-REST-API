package model

import "go.mongodb.org/mongo-driver/mongo"

type Users interface {
	Get(theUser User) (bool, error)
	All() ([]User, error)
	Create(theUser User) error
	Update(theUserId string, theUser User) (User, error)
	Delete(theUserId string) (User, error)
	DeleteAll() error
}

type userDB struct {
	collection *mongo.Collection
}

func CreateUserRepo(coll *mongo.Collection) Users {
	return &userDB{
		collection: coll,
	}
}

func (u *userDB) Get(theUser User) (bool, error)                      {}
func (u *userDB) All() ([]User, error)                                {}
func (u *userDB) Create(theUser User) error                           {}
func (u *userDB) Update(theUserId string, theUser User) (User, error) {}
func (u *userDB) Delete(theUserId string) (User, error)               {}
func (u *userDB) DeleteAll() error                                    {}
