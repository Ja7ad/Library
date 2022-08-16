package managers

import (
	"context"
	"github.com/Ja7ad/library/server/db/collection"
	"github.com/Ja7ad/library/server/global"
	"github.com/Ja7ad/library/server/internal/models"
	"github.com/Ja7ad/library/server/internal/models/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(ctx context.Context) ([]*schema.User, error) {
	users, err := models.GetAll[schema.User](ctx, collection.UserCollection())
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindUser(ctx context.Context, userID primitive.ObjectID) (*schema.User, error) {
	user, err := models.GetWithId[schema.User](ctx, collection.UserCollection(), userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(ctx context.Context, firstName, lastName string, age int32) (*schema.User, error) {
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	user := &schema.User{
		Id:        primitive.NewObjectID(),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	if err := models.Insert[schema.User](sessCtx, collection.UserCollection(), user); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(ctx context.Context, userID primitive.ObjectID, firstName, lastName string, age int32) (*schema.User, error) {
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	user, err := models.GetWithId[schema.User](sessCtx, collection.UserCollection(), userID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.Age = age

	if err := models.Update[schema.User](sessCtx, collection.UserCollection(), user.Id, user); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return err
	}

	user, err := models.GetWithId[schema.User](sessCtx, collection.UserCollection(), userID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := models.Delete[schema.User](sessCtx, collection.UserCollection(), user.Id); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return err
	}
	return nil
}

func ReserveBook(ctx context.Context, userID primitive.ObjectID, bookIDs ...primitive.ObjectID) error {
	sessCtx, err := global.Database.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.Database.StartTransaction(sessCtx); err != nil {
		return err
	}

	user, err := models.GetWithId[schema.User](sessCtx, collection.UserCollection(), userID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	user.BookIds = bookIDs
	if err := models.Update[schema.User](sessCtx, collection.UserCollection(), user.Id, user); err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := sessCtx.CommitTransaction(ctx); err != nil {
		return err
	}
	return nil
}
