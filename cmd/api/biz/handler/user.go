package handler

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/cmd/api/biz/rpc"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Register 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func Register(ctx context.Context, c *app.RequestContext) {
	var req user.CreateUserRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &user.CreateUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.CreateUser(ctx, &user.CreateUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &user.CreateUserResponse{StatusCode: Err.ErrCode, StatusMsg: &Err.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// Login 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func Login(ctx context.Context, c *app.RequestContext) {
	var req user.CheckUserRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		SetResponse(c, &user.CheckUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.CheckUse(context.Background(), &user.CheckUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		SetResponse(c, &user.CheckUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// GetUserInfo 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req user.GetUserRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		SetResponse(c, &user.GetUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.GetUserInfo(ctx, &user.GetUserRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		SetResponse(c, &user.GetUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}
