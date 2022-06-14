package models

import (
	"context"
	"github.com/Ja7ad/library/server/global"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Name          string             `bson:"name" json:"name"`
	PublisherId   primitive.ObjectID `bson:"publisher_id" json:"publisher_id"`
	PublisherName string             `bson:"publisher_name,omitempty" json:"publisher_name"`
	UserId        primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
}

func getCollection() *mongo.Collection {
	return global.BookClient.GetDatabase("library").Collection("book")
}

func GetBookByID(ctx context.Context, bookID primitive.ObjectID) (*Book, error) {
	book := &Book{}
	if err := getCollection().FindOne(ctx, bson.M{"_id": bookID}).Decode(book); err != nil {
		return nil, err
	}
	return book, nil
}

func GetBooks(ctx context.Context) ([]*Book, error) {
	books := []*Book{}
	cursor, err := getCollection().Aggregate(ctx, bookAggregatePipeline(), nil)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &books); err != nil {
		return nil, err
	}
	return books, nil
}

func FindBook(ctx context.Context, name, publisherName string, bookID, publisherID primitive.ObjectID) (*Book, error) {
	book := &Book{}
	filter := bson.M{}
	if len(name) != 0 {
		filter["name"] = name
	}
	if len(publisherName) != 0 {
		filter["publisher_name"] = publisherName
	}
	if !bookID.IsZero() {
		filter["_id"] = bookID
	}
	if !publisherID.IsZero() {
		filter["publisher_id"] = publisherID
	}
	stage := bson.D{{"$match", filter}}
	cursor, err := getCollection().Aggregate(ctx, bookAggregatePipeline(stage))
	if err != nil {
		return nil, err
	}
	if err := cursor.Decode(book); err != nil {
		return nil, err
	}
	return book, nil
}

func (b *Book) Insert(ctx context.Context) error {
	if _, err := getCollection().InsertOne(ctx, b); err != nil {
		return err
	}
	return nil
}

func (b *Book) Update(ctx context.Context) error {
	if _, err := getCollection().ReplaceOne(ctx, bson.M{"_id": b.Id}, b); err != nil {
		return err
	}
	return nil
}

func (b *Book) Delete(ctx context.Context) error {
	if _, err := getCollection().DeleteOne(ctx, bson.M{"_id": b.Id}); err != nil {
		return err
	}
	return nil
}

func bookAggregatePipeline(stages ...bson.D) mongo.Pipeline {
	pipeline := mongo.Pipeline{
		bson.D{{"$lookup", bson.M{
			"from":         "publisher",
			"localField":   "publisher_id",
			"foreignField": "_id",
			"as":           "pub",
		}}},
		bson.D{{"$unwind", bson.M{
			"path":                       "$pub",
			"preserveNullAndEmptyArrays": true,
		}}},
		bson.D{{"$addFields", bson.M{
			"publisher_name": "$pub.name",
		}}},
		bson.D{{"$project", bson.M{
			"publisher_id": 0,
		}}},
	}
	if len(stages) != 0 {
		for _, stage := range stages {
			pipeline = append(pipeline, stage)
		}
	}
	return pipeline
}
