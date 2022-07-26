package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type Connector interface {
	GetClient() *mongo.Client
	NewSession(context.Context, ...*options.SessionOptions) (mongo.SessionContext, error)
	StartTransaction(mongo.SessionContext) error
	SetDatabase(string)
	GetDatabase() *mongo.Database
	GetCollection(string) *mongo.Collection
}

type MongoClient struct {
	client   *mongo.Client
	database *mongo.Database
}

var _ Connector = (*MongoClient)(nil)

func NewMongo(ctx context.Context, uri string, opts ...*options.ClientOptions) (Connector, error) {
	opts = append(opts, options.Client().ApplyURI(uri))
	cli, err := mongo.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	if err := cli.Connect(ctx); err != nil {
		return nil, err
	}
	if err := cli.Ping(ctx, nil); err != nil {
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

func (m *MongoClient) StartTransaction(sessCtx mongo.SessionContext) error {
	if err := sessCtx.StartTransaction(options.Transaction().
		SetReadConcern(readconcern.Snapshot()).
		SetWriteConcern(writeconcern.New(writeconcern.WMajority()))); err != nil {
		return err
	}
	return nil
}

func (m *MongoClient) SetDatabase(dbName string) {
	m.database = m.client.Database(dbName)
}

func (m *MongoClient) GetDatabase() *mongo.Database {
	return m.database
}

func (m *MongoClient) GetCollection(collectionName string) *mongo.Collection {
	return m.database.Collection(collectionName)
}
