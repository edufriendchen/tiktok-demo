package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/edufriendchen/tiktok-demo/cmd/api/biz/rpc"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/relation"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
)

// RelationAction 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func RelationAction(ctx context.Context, c *app.RequestContext) {
	var req relation.ActionRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &user.CreateUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.ActionRelation(ctx, &req)
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &user.CreateUserResponse{StatusCode: Err.ErrCode, StatusMsg: &Err.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// MGetFollowList 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func MGetFollowList(ctx context.Context, c *app.RequestContext) {
	var req relation.FollowRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		SetResponse(c, &user.CheckUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.MGetFollowList(context.Background(), &req)
	if err != nil {
		SetResponse(c, &user.CheckUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// MGetFollowerList 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func MGetFollowerList(ctx context.Context, c *app.RequestContext) {
	var req relation.FollowerRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		SetResponse(c, &user.MGetUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.MGetFollowerList(ctx, &req)
	if err != nil {
		SetResponse(c, &user.MGetUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// MGetUserInfo 注册用户操作 的上下文至 User 服务的 RPC 客户端, 并获取相应的响应.
func MGetFriendList(ctx context.Context, c *app.RequestContext) {
	var req relation.FriendRequest
	err := c.BindAndValidate(&req)
	if err != nil {
		SetResponse(c, &user.MGetUserResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.MGetFriendList(ctx, &req)
	if err != nil {
		SetResponse(c, &user.MGetUserResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}
