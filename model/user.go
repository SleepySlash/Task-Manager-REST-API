package model

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	Get(theUser User) (bool, error)
	All() ([]User, error)
	Create(theUser User) error
	Update(theUserId string, theUser User) (int, error)
	Delete(theUserId string) error
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

func (u *userDB) Get(theUser User) (bool, error) {
	var res User
	filter := bson.D{{Key: "username", Value: theUser.Username}, {Key: "password", Value: theUser.password}}
	err := u.collection.FindOne(context.TODO(), filter).Decode(&res)
	if err != nil {
		return false, err
	}
	log.Println("found user in the db", res)
	return true, nil
}

func (u *userDB) All() ([]User, error) {
	cursor, err := u.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var result []User
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userDB) Create(theUser User) error {
	res, err := u.collection.InsertOne(context.TODO(), theUser)
	if err != nil {
		return err
	}
	log.Println("inserted user into user db", res)
	return nil
}

func (u *userDB) Update(theUserId string, theUser User) (int, error) {
	theUser.UserId = theUserId
	filter := bson.D{{Key: "user_id", Value: theUserId}}
	update := bson.D{{Key: "$set", Value: theUser}}
	updateRes, err := u.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 0, err
	}
	log.Println("updated the user", updateRes)
	return int(updateRes.ModifiedCount), nil
}

func (u *userDB) Delete(theUserId string) error {
	res, err := u.collection.DeleteOne(context.TODO(), bson.D{{Key: "user_id", Value: theUserId}})
	if err != nil {
		return err
	}
	log.Println("deleted user from db", res.DeletedCount)
	return nil
}

func (u *userDB) DeleteAll() error {
	res, err := u.collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		return err
	}
	log.Println("delted user from db", res.DeletedCount)
	return nil
}
