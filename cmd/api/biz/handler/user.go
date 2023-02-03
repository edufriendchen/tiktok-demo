package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/edufriendchen/tiktok-demo/cmd/api/biz/rpc"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
)

// Register 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func Register(ctx context.Context, c *app.RequestContext) {
	var req user.CreateUserRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &user.CreateUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.CreateUser(ctx, &req)
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
	resp, err := rpc.CheckUse(context.Background(), &req)
	if err != nil {
		SetResponse(c, &user.CheckUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// MGetUserInfo 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func MGetUserInfo(ctx context.Context, c *app.RequestContext) {
	var req user.MGetUserRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		SetResponse(c, &user.MGetUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.MGetUserInfo(ctx, &req)
	if err != nil {
		SetResponse(c, &user.MGetUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	resp.User.Avatar = "https://img1.baidu.com/it/u=1459539381,1684299919&fm=253&app=138&size=w931&n=0&f=JPEG&fmt=auto?sec=1675530000&t=264b8cfbd62ce23ee0d0a557091cc72d"
	SetResponse(c, resp)
	return
}
