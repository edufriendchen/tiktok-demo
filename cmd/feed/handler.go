package main

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/cmd/feed/service"
	feed "github.com/edufriendchen/tiktok-demo/kitex_gen/feed"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/edufriendchen/tiktok-demo/pkg/global"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// MGetFeedList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) MGetFeedList(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	resp = new(feed.FeedResponse)
	if err = req.IsValid(); err != nil {
		resp = &feed.FeedResponse{StatusCode: errno.ParamErr.ErrCode, StatusMsg: &errno.ParamErr.ErrMsg}
		return resp, err
	}
	list, err := service.NewMGetFeedListService(ctx, global.Neo4jDriver).MGetFeedList(req, consts.Limit)
	if err != nil {
		resp = &feed.FeedResponse{StatusCode: errno.ServiceErr.ErrCode, StatusMsg: &errno.ServiceErr.ErrMsg}
		return resp, err
	}
	resp = &feed.FeedResponse{StatusCode: errno.Success.ErrCode, StatusMsg: &errno.Success.ErrMsg}
	resp.VideoList = list
	return resp, nil
}
