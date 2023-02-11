package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/edufriendchen/tiktok-demo/cmd/api/biz/rpc"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/message"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
)

// ActionMessage
func ActionMessage(ctx context.Context, c *app.RequestContext) {
	fmt.Println("进入:")
	var req message.ActionRequest
	err := c.BindAndValidate(&req)
	fmt.Println("params:", req)
	if err != nil {
		c.JSON(consts.StatusOK, &message.ActionResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.ActionMessage(ctx, &req)
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &message.ActionResponse{StatusCode: Err.ErrCode, StatusMsg: &Err.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}

// MGetMessageList
func MGetMessageList(ctx context.Context, c *app.RequestContext) {
	var req message.ChatRequest
	err := c.BindAndValidate(&req)
	fmt.Println("params:", req)
	if err != nil {
		c.JSON(consts.StatusOK, &message.ChatResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.MGetChatMessage(ctx, &req)
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &message.ChatResponse{StatusCode: Err.ErrCode, StatusMsg: &Err.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}
