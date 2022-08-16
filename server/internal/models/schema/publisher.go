package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type Publisher struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
}
