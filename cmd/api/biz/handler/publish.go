package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/edufriendchen/tiktok-demo/cmd/api/biz/rpc"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/publish"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
)

// PublishAction 注册用户操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func PublishAction(ctx context.Context, c *app.RequestContext) {
	var req publish.ActionRequest
	fmt.Println("req:", req)
	err := c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusOK, &publish.ActionResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.ActionPublish(ctx, &req)
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &publish.ActionResponse{StatusCode: Err.ErrCode, StatusMsg: &Err.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// MGetPublishList 注册用户操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func MGetPublishList(ctx context.Context, c *app.RequestContext) {
	var req publish.PublishRequest
	err := c.BindAndValidate(&req)
	fmt.Println("req:", req)
	if err != nil {
		SetResponse(c, &publish.PublishResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.MGetPublishList(ctx, &req)
	if err != nil {
		SetResponse(c, &publish.PublishResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}
