package rpc

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/message"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/message/messageservice"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/edufriendchen/tiktok-demo/pkg/initialize"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/resolver"
)

var messageClient messageservice.Client

func initMessage() {
	cli, err := initialize.InitNacos()
	c, err := messageservice.NewClient(
		consts.MessageServiceName,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

// ActionRelation
func ActionMessage(ctx context.Context, req *message.ActionRequest) (*message.ActionResponse, error) {
	resp, err := messageClient.ActionMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

// ChatMessage
func MGetChatMessage(ctx context.Context, req *message.ChatRequest) (*message.ChatResponse, error) {
	resp, err := messageClient.ChatMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	fmt.Println("Return:", resp)
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
