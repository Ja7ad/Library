package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transactor interface {
	GetClient() *mongo.Client
	GetDatabase(string) *mongo.Database
	NewSession(context.Context, ...*options.SessionOptions) (mongo.SessionContext, error)
}

type MongoClient struct {
	client *mongo.Client
}

func NewMongo(uri string) (Transactor, error) {
	cli, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := cli.Connect(context.Background()); err != nil {
		return nil, err
	}
	return &MongoClient{client: cli}, nil
}

func (m *MongoClient) GetClient() *mongo.Client {
	return m.client
}

func (m *MongoClient) NewSession(ctx context.Context, opts ...*options.SessionOptions) (mongo.SessionContext, error) {
	sess, err := m.client.StartSession(opts...)
	if err != nil {
		return nil, err
	}
	return mongo.NewSessionContext(ctx, sess), nil
}

func (m *MongoClient) GetDatabase(dbName string) *mongo.Database {
	return m.client.Database(dbName)
}
