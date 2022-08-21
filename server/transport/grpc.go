package transport

import (
	"context"
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"github.com/Ja7ad/library/server/internal/service"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
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

	descriptor := make(map[string]*desc.MethodDescriptor)
	srv := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		recoveryMiddleware(),
		methodDescriptors(descriptor),
		jwtMiddleware(),
	))
	library.RegisterBookServiceServer(srv, &service.LibraryServer{})
	library.RegisterUserServiceServer(srv, &service.UserServer{})
	grpc_health_v1.RegisterHealthServer(srv, &service.HealthyServer{})
	reflection.Register(srv)

	serviceDesc, err := grpcreflect.LoadServiceDescriptors(srv)
	if err != nil {
		return nil, err
	}

	for _, d := range serviceDesc {
		for _, md := range d.GetMethods() {
			methodName := fmt.Sprintf("/%s/%s", d.GetFullyQualifiedName(), md.GetName())
			descriptor[methodName] = md
		}
	}

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

func recoveryMiddleware() grpc.UnaryServerInterceptor {
	rec := func(p interface{}) (err error) {
		err = status.Errorf(codes.Unknown, "%v", p)
		log.Printf("panic triggered: Error %v led to gRPC server recovery", err)
		return
	}
	opts := []grpc_recovery.Option{
		grpc_recovery.WithRecoveryHandler(rec),
	}
	return grpc_recovery.UnaryServerInterceptor(opts...)

}

func methodDescriptors(descriptors map[string]*desc.MethodDescriptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md := descriptors[info.FullMethod]
		ctx = metadata.AppendToOutgoingContext(context.WithValue(ctx, "desc", md))
		return handler(ctx, req)
	}
}

func jwtMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		d, ok := ctx.Value("desc").(*desc.MethodDescriptor)
		if !ok {
			return handler(ctx, req)
		}
		option := proto.GetExtension(d.GetMethodOptions(), library.E_Permission)
		permission, ok := option.(*library.Permission)
		if !ok {
			return handler(ctx, req)
		}
		fmt.Printf("Method %s Permission %v", info.FullMethod, permission.PermissionCodes)
		return handler(ctx, req)
	}
}
