package models

import (
	"context"
	"github.com/Ja7ad/library/server/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var colUser = global.UserClient.GetDatabase("user").Collection("user")

type User struct {
	Id        primitive.ObjectID   `bson:"_id" json:"id"`
	FirstName string               `bson:"first_name" json:"first_name"`
	LastName  string               `bson:"last_name" json:"last_name"`
	Age       int32                `bson:"age,omitempty" json:"age"`
	BookIds   []primitive.ObjectID `bson:"book_ids,omitempty" json:"book_ids"`
}

func GetUserById(ctx context.Context, userID primitive.ObjectID) (*User, error) {
	user := &User{}
	if err := colUser.FindOne(ctx, bson.M{"_id": userID}).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func FindUser(ctx context.Context) {

}
