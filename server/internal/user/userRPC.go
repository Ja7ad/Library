package user

import (
	"context"
	"github.com/Ja7ad/library/proto/protoModel/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	user.UnimplementedUserServiceServer
}

func (*UserServer) AddUser(ctx context.Context, request *user.AddUserRequest) (*user.User, error) {
	userModel, err := AddUser(ctx, request.FirstName, request.LastName, request.Age)
	if err != nil {
		return nil, err
	}
	return &user.User{
		Id:        userModel.Id.Hex(),
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Age:       userModel.Age,
	}, nil
}

func (*UserServer) UpdateUser(ctx context.Context, request *user.UpdateUserRequest) (*user.User, error) {
	userID, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	userModel, err := UpdateUser(ctx, userID, request.FirstName, request.LastName, request.Age)
	if err != nil {
		return nil, err
	}
	return &user.User{
		Id:        userModel.Id.Hex(),
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Age:       userModel.Age,
	}, nil
}

func (*UserServer) ReserveBook(ctx context.Context, request *user.ReserveUserBookRequest) (*emptypb.Empty, error) {
	bookIDs := []primitive.ObjectID{}
	for _, b := range request.BookId {
		bookID, err := primitive.ObjectIDFromHex(b)
		if err != nil {
			return nil, err
		}
		bookIDs = append(bookIDs, bookID)
	}
	userID, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	if err := ReserveBook(ctx, userID, bookIDs...); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (*UserServer) DeleteUser(ctx context.Context, request *user.DeleteUserRequest) (*emptypb.Empty, error) {
	userID, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	if err := DeleteUser(ctx, userID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
