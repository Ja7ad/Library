package user

import (
	"context"
	"fmt"
	"github.com/Ja7ad/library/proto/protoModel/user"
	bookDB "github.com/Ja7ad/library/server/db"
	"github.com/Ja7ad/library/server/global"
	"github.com/ory/dockertest/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const (
	bufSize = 1024 * 1024
)

var (
	lis      *bufconn.Listener
	pool     *dockertest.Pool
	resource *dockertest.Resource
	err      error
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

	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatal(err)
	}

	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=root",
		"MONGO_INITDB_ROOT_PASSWORD=password",
	}

	resource, err = pool.Run("mongo", "5.0", environmentVariables)
	if err != nil {
		log.Fatal(err)
	}

	if err = pool.Retry(func() error {
		ctx := context.TODO()
		db, err := bookDB.NewMongo(ctx, fmt.Sprintf("mongodb://root:password@localhost:%s", resource.GetPort("27017/tcp")))
		if err != nil {
			return err
		}
		global.UserClient = db
		global.UserClient.SetDatabase("user")
		return nil
	}); err != nil {
		log.Fatal(err)
	}
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

func TestCleanDatabase(t *testing.T) {
	if err = pool.Purge(resource); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}
}
