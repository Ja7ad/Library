package models

import (
	"context"
	"github.com/Ja7ad/library/server/internal/models/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Library interface {
	schema.Book | schema.Publisher | schema.User
}

func GetWithId[T Library](ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (*T, error) {
	var data *T
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetWithName[T Library](ctx context.Context, collection *mongo.Collection, name string) (*T, error) {
	var data *T
	if err := collection.FindOne(ctx, bson.M{"name": name}).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func GetAll[T Library](ctx context.Context, collection *mongo.Collection) ([]*T, error) {
	var data []*T
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func Insert[T Library](ctx context.Context, collection *mongo.Collection, data *T) error {
	if _, err := collection.InsertOne(ctx, data); err != nil {
		return err
	}
	return nil
}

func Update[T Library](ctx context.Context, collection *mongo.Collection, id primitive.ObjectID, data *T) error {
	if _, err := collection.UpdateOne(ctx, bson.M{"_id": id}, data); err != nil {
		return err
	}
	return nil
}

func Delete[T Library](ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) error {
	if _, err := collection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return err
	}
	return nil
}
