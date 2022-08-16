package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Name          string             `bson:"name" json:"name"`
	PublisherId   primitive.ObjectID `bson:"publisher_id" json:"publisher_id"`
	PublisherName string             `bson:"publisher_name,omitempty" json:"publisher_name"`
	UserId        primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
}
