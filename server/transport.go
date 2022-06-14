package main

import (
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitServer(addr, port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		return err
	} else {
		log.Println("server ran on :3345")
	}
	srv := grpc.NewServer()
	library.RegisterLibraryServiceServer(srv, &LibraryServer{})

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
