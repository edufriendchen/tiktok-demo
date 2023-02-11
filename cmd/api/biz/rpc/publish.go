package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/publish"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/publish/publishservice"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/edufriendchen/tiktok-demo/pkg/initialize"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

var publishClient publishservice.Client

func initPublish() {
	cli, err := initialize.InitNacos()
	c, err := publishservice.NewClient(
		consts.PushlishServiceName,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

// ActionPublish
func ActionPublish(ctx context.Context, req *publish.ActionRequest) (*publish.ActionResponse, error) {
	resp, err := publishClient.ActionPulish(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

// MGetPublishList
func MGetPublishList(ctx context.Context, req *publish.PublishRequest) (*publish.PublishResponse, error) {
	resp, err := publishClient.MGetPublishList(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
