package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID   `bson:"_id" json:"id"`
	FirstName string               `bson:"first_name" json:"first_name"`
	LastName  string               `bson:"last_name" json:"last_name"`
	Age       int32                `bson:"age,omitempty" json:"age"`
	BookIds   []primitive.ObjectID `bson:"book_ids,omitempty" json:"book_ids"`
}
