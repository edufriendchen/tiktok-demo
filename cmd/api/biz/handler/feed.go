package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/edufriendchen/tiktok-demo/cmd/api/biz/rpc"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/feed"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
)

// MGetFeedList
func MGetFeedList(ctx context.Context, c *app.RequestContext) {
	fmt.Println("进入:")
	var req feed.FeedRequest
	err := c.BindAndValidate(&req)
	fmt.Println("params:", req)
	if err != nil {
		c.JSON(consts.StatusOK, &feed.FeedResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg})
		return
	}
	resp, err := rpc.MGetFeedList(ctx, &req)
	if err != nil {
		Err := errno.ConvertErr(err)
		c.JSON(consts.StatusOK, &feed.FeedResponse{StatusCode: Err.ErrCode, StatusMsg: &Err.ErrMsg})
		return
	}
	SetResponse(c, resp)
	return
}
