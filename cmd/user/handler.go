package main

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/cmd/user/global"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"

	service "github.com/edufriendchen/tiktok-demo/cmd/user/service"

	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/jwt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)
	if err = req.IsValid(); err != nil {
		resp = &user.CreateUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg}
		return resp, err
	}
	userid, err := service.NewCreateUserService(ctx).CreateUserNode(ctx, global.Neo4jSession, req)
	if err != nil {
		resp = &user.CreateUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg}
		return resp, err
	}
	resp = &user.CreateUserResponse{StatusCode: errno.Success.ErrCode, StatusMsg: &errno.Success.ErrMsg, UserId: userid, Token: ""}
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)
	if err = req.IsValid(); err != nil {
		return resp, err
	}
	userid, err := service.NewCheckUserService(ctx).Login(ctx, global.Neo4jSession, req)
	if err != nil {
		return resp, err
	}
	token, err := global.Jwt.CreateToken(jwt.CustomClaims{
		Id: userid,
	})
	resp = &user.CheckUserResponse{StatusCode: errno.Success.ErrCode, StatusMsg: &errno.Success.ErrMsg, UserId: userid, Token: token}
	return resp, nil
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	resp = new(user.GetUserResponse)
	if err = req.IsValid(); err != nil {
		return resp, err
	}
	user, err := service.NewGetUserService(ctx).GetUser(ctx, global.Neo4jSession, req)
	if err != nil {
		return resp, err
	}
	resp.User = user
	return resp, nil
}
