package collection

import (
	"github.com/Ja7ad/library/server/global"
	"go.mongodb.org/mongo-driver/mongo"
)

func BookCollection() *mongo.Collection {
	return global.Database.GetCollection("book")
}

func PublisherCollection() *mongo.Collection {
	return global.Database.GetCollection("publisher")
}

func UserCollection() *mongo.Collection {
	return global.Database.GetCollection("user")
}
