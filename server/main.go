package main

import (
	"context"
	bookDB "github.com/Ja7ad/library/server/db"
	"github.com/Ja7ad/library/server/global"
	"github.com/Ja7ad/library/server/transport/grpc"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	bookURI := os.Getenv("MONGO_BOOK_URI")
	if bookURI == "" {
		log.Fatal("You must set your 'MONGO_BOOK_URI' environmental variable")
	}

	userURI := os.Getenv("MONGO_USER_URI")
	if userURI == "" {
		log.Fatal("You must set your 'MONGO_USER_URI' environmental variable")
	}

	transBook, err := bookDB.NewMongo(context.Background(), bookURI)
	if err != nil {
		log.Fatal(err)
	}

	transUser, err := bookDB.NewMongo(context.Background(), userURI)
	if err != nil {
		log.Fatal(err)
	}

	global.BookClient = transBook
	global.BookClient.SetDatabase("book")

	global.UserClient = transUser
	global.UserClient.SetDatabase("user")
}

func main() {
	if err := grpc.InitServer("localhost", "3345"); err != nil {
		log.Fatal(err)
	}
}
