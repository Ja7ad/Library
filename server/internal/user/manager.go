package user

import (
	"context"
	"github.com/Ja7ad/library/server/global"
	"github.com/Ja7ad/library/server/internal/user/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(ctx context.Context) ([]*models.User, error) {
	users, err := models.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FindUser(ctx context.Context, userID primitive.ObjectID) (*models.User, error) {
	user, err := models.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AddUser(ctx context.Context, firstName, lastName string, age int32) (*models.User, error) {
	sessCtx, err := global.UserClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.UserClient.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	user := &models.User{
		Id:        primitive.NewObjectID(),
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	if err := user.Insert(sessCtx); err != nil {
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

func UpdateUser(ctx context.Context, userID primitive.ObjectID, firstName, lastName string, age int32) (*models.User, error) {
	sessCtx, err := global.UserClient.NewSession(ctx)
	if err != nil {
		return nil, err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.UserClient.StartTransaction(sessCtx); err != nil {
		return nil, err
	}

	user, err := models.GetUserById(sessCtx, userID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return nil, err
		}
		return nil, err
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.Age = age

	if err := user.Update(sessCtx); err != nil {
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
	sessCtx, err := global.UserClient.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.UserClient.StartTransaction(sessCtx); err != nil {
		return err
	}

	user, err := models.GetUserById(sessCtx, userID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	if err := user.Delete(sessCtx); err != nil {
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
	sessCtx, err := global.UserClient.NewSession(ctx)
	if err != nil {
		return err
	}
	defer sessCtx.EndSession(ctx)

	if err := global.UserClient.StartTransaction(sessCtx); err != nil {
		return err
	}

	user, err := models.GetUserById(sessCtx, userID)
	if err != nil {
		if err := sessCtx.AbortTransaction(ctx); err != nil {
			return err
		}
		return err
	}

	user.BookIds = bookIDs
	if err := user.Update(sessCtx); err != nil {
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
