package main

import (
	"context"
	bookDB "github.com/Ja7ad/library/server/db"
	"github.com/Ja7ad/library/server/global"
	"github.com/Ja7ad/library/server/transport"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	libraryURI := os.Getenv("MONGO_LIBRARY_URI")
	if libraryURI == "" {
		log.Fatal("You must set your 'MONGO_LIBRARY_URI' environmental variable")
	}

	transLibrary, err := bookDB.NewMongo(context.Background(), libraryURI)
	if err != nil {
		log.Fatal(err)
	}

	global.Database = transLibrary
	global.Database.SetDatabase("library")
}

func main() {
	clientCon, err := transport.InitGrpcService(os.Getenv("SERVER_GRPC_ADDRESS"), os.Getenv("SERVER_GRPC_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	if err := transport.InitRestService(os.Getenv("SERVER_HTTP_ADDRESS"), os.Getenv("SERVER_HTTP_PORT"), clientCon); err != nil {
		log.Fatal(err)
	}
}
