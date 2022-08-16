package transport

import (
	"context"
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"github.com/Ja7ad/library/server/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func InitGrpcService(addr, port string) (*grpc.ClientConn, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		return nil, err
	} else {
		log.Printf("grpc server ran on %s:%s", addr, port)
	}
	srv := grpc.NewServer()
	library.RegisterBookServiceServer(srv, &service.LibraryServer{})
	library.RegisterUserServiceServer(srv, &service.UserServer{})
	grpc_health_v1.RegisterHealthServer(srv, &service.HealthyServer{})
	reflection.Register(srv)
	go func() {
		log.Fatalln(srv.Serve(listener))
	}()

	maxMsgSize := 1024 * 1024 * 20
	return grpc.DialContext(
		context.Background(),
		fmt.Sprintf("%s:%s", addr, port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize), grpc.MaxCallSendMsgSize(maxMsgSize)),
	)
}
