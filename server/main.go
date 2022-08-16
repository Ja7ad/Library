package main

import (
	"context"
	"expvar"
	"fmt"
	bookDB "github.com/Ja7ad/library/server/db"
	"github.com/Ja7ad/library/server/global"
	"github.com/Ja7ad/library/server/transport"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/http/pprof"
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
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("DEBUG_PORT")), pprofService()); err != nil {
			log.Fatal(err)
		}
	}()

	clientCon, err := transport.InitGrpcService(os.Getenv("SERVER_GRPC_ADDRESS"), os.Getenv("SERVER_GRPC_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	if err := transport.InitRestService(os.Getenv("SERVER_HTTP_ADDRESS"), os.Getenv("SERVER_HTTP_PORT"), clientCon); err != nil {
		log.Fatal(err)
	}
}

func pprofService() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/vars", expvar.Handler())
	return mux
}
