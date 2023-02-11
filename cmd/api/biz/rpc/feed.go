package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/feed"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/feed/feedservice"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/edufriendchen/tiktok-demo/pkg/initialize"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

var feedClient feedservice.Client

func initFeed() {
	cli, err := initialize.InitNacos()
	c, err := feedservice.NewClient(
		consts.FeedServiceName,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}

// MGetFeedList
func MGetFeedList(ctx context.Context, req *feed.FeedRequest) (*feed.FeedResponse, error) {
	resp, err := feedClient.MGetFeedList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
