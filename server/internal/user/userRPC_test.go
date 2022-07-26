package user

import (
	"context"
	"github.com/Ja7ad/library/proto/protoModel/user"
	bookDB "github.com/Ja7ad/library/server/db"
	"github.com/Ja7ad/library/server/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"
)

const (
	bufSize = 1024 * 1024
)

var (
	lis      *bufconn.Listener
	mongoURI = os.Getenv("MONGO_USER_URI")
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &UserServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	client, err := bookDB.NewMongo(mongoURI)
	if err != nil {
		log.Fatal(err)
	}

	global.UserClient = client
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddUser(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := user.NewUserServiceClient(conn)
	users := []*user.AddUserRequest{
		{FirstName: "FirstA", LastName: "LastA", Age: 20},
		{FirstName: "FirstB", LastName: "LastB", Age: 21},
		{FirstName: "FirstC", LastName: "LastC", Age: 22},
		{FirstName: "FirstD", LastName: "LastD", Age: 23},
	}
	for _, request := range users {
		resp, err := client.AddUser(ctx, request)
		if err != nil {
			t.Error(err)
		}
		t.Log(resp)
	}

}
