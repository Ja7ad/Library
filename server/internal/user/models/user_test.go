package models

import (
	"context"
	"fmt"
	bookDB "github.com/Ja7ad/library/server/db"
	"github.com/Ja7ad/library/server/global"
	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var (
	pool     *dockertest.Pool
	resource *dockertest.Resource
	err      error
)

func TestClient(t *testing.T) {
	pool, err = dockertest.NewPool("")
	if err != nil {
		t.Error(err)
	}

	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=root",
		"MONGO_INITDB_ROOT_PASSWORD=password",
	}

	resource, err = pool.Run("mongo", "5.0", environmentVariables)
	if err != nil {
		t.Error(err)
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
		t.Errorf("Could not connect to docker: %s", err)
	}

}

func TestAddUser(t *testing.T) {
	users := []*User{
		{
			Id:        primitive.NewObjectID(),
			FirstName: "FirstA",
			LastName:  "LastA",
			Age:       20,
			BookIds:   []primitive.ObjectID{primitive.NewObjectID(), primitive.NewObjectID()},
		},
		{
			Id:        primitive.NewObjectID(),
			FirstName: "FirstB",
			LastName:  "LastB",
			Age:       21,
			BookIds:   []primitive.ObjectID{primitive.NewObjectID()},
		},
		{
			Id:        primitive.NewObjectID(),
			FirstName: "FirstC",
			LastName:  "LastC",
			Age:       22,
			BookIds:   []primitive.ObjectID{},
		},
		{
			Id:        primitive.NewObjectID(),
			FirstName: "FirstD",
			LastName:  "LastD",
			BookIds:   []primitive.ObjectID{},
		},
		{
			Id:        primitive.NewObjectID(),
			FirstName: "FirstD",
			LastName:  "LastD",
			Age:       23,
		},
	}
	ctx := context.TODO()
	for _, user := range users {
		if err := user.Insert(ctx); err != nil {
			t.Error(err)
		}
		t.Log(user)
	}
}

func TestGetUsers(t *testing.T) {
	ctx := context.TODO()
	users, err := GetUsers(ctx)
	if err != nil {
		t.Error(err)
	}
	for _, user := range users {
		t.Log(user)
	}
}

func TestCleanDatabase(t *testing.T) {
	if err = pool.Purge(resource); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}
}
