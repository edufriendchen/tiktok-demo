package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user/userservice"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

var userClient userservice.Client

func initUser() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "info",
		Username:            "nacos",
		Password:            "nacos",
	}
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

// CheckUser check user info
func CheckUse(ctx context.Context, req *user.CheckUserRequest) (*user.CheckUserResponse, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return resp, err
	}
	if resp.StatusCode != 0 {
		return resp, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}

// Login rpc to user service GetUserInfo
func GetUserInfo(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	resp, err := userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp, nil
}
