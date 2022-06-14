package user

import "github.com/Ja7ad/library/proto/protoModel/user"

type UserServer struct {
	user.UnimplementedUserServiceServer
}
