package models

import (
	"context"
	"github.com/Ja7ad/library/server/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id        primitive.ObjectID   `bson:"_id" json:"id"`
	FirstName string               `bson:"first_name" json:"first_name"`
	LastName  string               `bson:"last_name" json:"last_name"`
	Age       int32                `bson:"age,omitempty" json:"age"`
	BookIds   []primitive.ObjectID `bson:"book_ids,omitempty" json:"book_ids"`
}

func getCollection() *mongo.Collection {
	return global.UserClient.GetDatabase("user").Collection("user")
}

func GetUserById(ctx context.Context, userID primitive.ObjectID) (*User, error) {
	user := &User{}
	if err := getCollection().FindOne(ctx, bson.M{"_id": userID}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func GetUsers(ctx context.Context) ([]*User, error) {
	users := []*User{}
	cursor, err := getCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) Insert(ctx context.Context) error {
	if _, err := getCollection().InsertOne(ctx, u); err != nil {
		return err
	}
	return nil
}

func (u *User) Update(ctx context.Context) error {
	if _, err := getCollection().ReplaceOne(ctx, bson.M{"_id": u.Id}, u); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(ctx context.Context) error {
	if _, err := getCollection().DeleteOne(ctx, bson.M{"_id": u.Id}); err != nil {
		return err
	}
	return nil
}
