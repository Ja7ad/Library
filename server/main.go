package main

import (
	"github.com/Ja7ad/library/server/book/db"
	"github.com/Ja7ad/library/server/book/global"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable")
	}

	trans, err := db.NewMongo(uri)
	if err != nil {
		log.Fatal(err)
	}
	global.Client = trans

	if err := InitServer("localhost", "3345"); err != nil {
		log.Fatal(err)
	}
}
