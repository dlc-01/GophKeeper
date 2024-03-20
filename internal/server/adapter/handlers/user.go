package handlers

import (
	"context"
	proto "github.com/dlc-01/GophKeeper/internal/general/proto/gen"
	"github.com/dlc-01/GophKeeper/internal/server/core/domain/models"
	"github.com/dlc-01/GophKeeper/internal/server/core/port"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	proto.UnimplementedUserServer
	user port.IUsersService
	auth port.IAuthService
}

func NewUserServer(user port.IUsersService, auth port.IAuthService) *UserServer {
	return &UserServer{
		user: user,
		auth: auth,
	}
}

func (s *UserServer) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	var resp proto.DeleteUserResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	user := models.User{
		ID: &userID,
	}

	_, err := s.user.Delete(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	var resp proto.UpdateUserResponse

	userID, ok := ctx.Value(UserIDKey).(uint64)
	if !ok {
		return nil, status.Error(codes.Aborted, "missing user_id")
	}

	newUser := models.User{
		ID:           &userID,
		Username:     req.User.GetLogin(),
		PasswordHash: req.User.GetPasswordHash(),
	}

	users, err := s.user.Update(ctx, &newUser)
	if err != nil {
		return nil, err
	}

	token, err := s.auth.Login(ctx, users)
	if err != nil {
		return nil, err
	}

	resp.Token = token

	return &resp, nil
}
