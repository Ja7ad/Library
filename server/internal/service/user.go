package service

import (
	"context"
	"github.com/Ja7ad/library/proto/protoModel/library"
	"github.com/Ja7ad/library/server/internal/managers"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	library.UnimplementedUserServiceServer
}

func (*UserServer) AddUser(ctx context.Context, request *library.AddUserRequest) (*library.User, error) {
	userModel, err := managers.AddUser(ctx, request.FirstName, request.LastName, request.Age)
	if err != nil {
		return nil, err
	}
	return &library.User{
		Id:        userModel.Id.Hex(),
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Age:       userModel.Age,
	}, nil
}

func (*UserServer) UpdateUser(ctx context.Context, request *library.UpdateUserRequest) (*library.User, error) {
	userID, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	userModel, err := managers.UpdateUser(ctx, userID, request.FirstName, request.LastName, request.Age)
	if err != nil {
		return nil, err
	}
	return &library.User{
		Id:        userModel.Id.Hex(),
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Age:       userModel.Age,
	}, nil
}

func (*UserServer) DeleteUser(ctx context.Context, request *library.DeleteUserRequest) (*emptypb.Empty, error) {
	userID, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return nil, err
	}
	if err := managers.DeleteUser(ctx, userID); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
