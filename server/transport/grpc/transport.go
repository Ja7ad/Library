package grpc

import (
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"github.com/Ja7ad/library/proto/protoModel/user"
	"github.com/Ja7ad/library/server/internal/book"
	userRPC "github.com/Ja7ad/library/server/internal/user"
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
	library.RegisterLibraryServiceServer(srv, &book.LibraryServer{})
	user.RegisterUserServiceServer(srv, &userRPC.UserServer{})

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
