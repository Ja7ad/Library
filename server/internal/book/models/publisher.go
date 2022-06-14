package models

import (
	"context"
	"github.com/Ja7ad/library/server/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Publisher struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
}

func getPubCollection() *mongo.Collection {
	return global.BookClient.GetDatabase("library").Collection("publisher")
}

func (b *Publisher) Insert(ctx context.Context) error {
	if _, err := getPubCollection().InsertOne(ctx, b); err != nil {
		return err
	}
	return nil
}

func GetPublisherByName(ctx context.Context, publisherName string) (*Publisher, error) {
	publisher := &Publisher{}
	if err := getPubCollection().FindOne(ctx, bson.M{"name": publisherName}).Decode(publisher); err != nil {
		return nil, err
	}
	return publisher, nil
}
