package transport

import (
	"context"
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"github.com/Ja7ad/library/server/internal/service"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"runtime/debug"
)

func InitGrpcService(addr, port string) (*grpc.ClientConn, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		return nil, err
	} else {
		log.Printf("grpc server ran on %s:%s", addr, port)
	}
	srv := grpcServer()
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

func grpcServer() *grpc.Server {
	return grpc.NewServer(middlewares())
}

func middlewares() grpc.ServerOption {
	rec := func(p interface{}) (err error) {
		err = status.Errorf(codes.Unknown, "%v", p)
		log.Printf("panic triggered: Error %v led to gRPC server recovery \n\n%s", err, string(debug.Stack()))
		return
	}

	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(rec),
	}

	return grpc_middleware.WithUnaryServerChain(
		grpc_recovery.UnaryServerInterceptor(opts...),
	)
}
