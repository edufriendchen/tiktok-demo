package main

import (
	"context"
	"log"
	"net"

	"github.com/edufriendchen/tiktok-demo/cmd/user/global"
	"github.com/edufriendchen/tiktok-demo/cmd/user/initialize"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/jwt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	api "github.com/edufriendchen/tiktok-demo/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

// User RPC Server 端配置初始化
func Init() {
	global.Jwt = jwt.NewJWT([]byte(consts.JWTSecretKey))
}

func main() {

	cli, err := initialize.InitNacos()
	if err != nil {
		panic(err)
	}

	driver, err := neo4j.NewDriverWithContext(consts.Neo4jUri, neo4j.BasicAuth("neo4j", "friendchen0429", ""))
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer driver.Close(ctx)
	global.Neo4jSession = driver.NewSession(ctx, neo4j.SessionConfig{})
	defer global.Neo4jSession.Close(ctx)

	addr, err := net.ResolveTCPAddr(consts.TCP, consts.UserServiceAddr)
	if err != nil {
		panic(err)
	}
	Init()

	svr := api.NewServer(new(UserServiceImpl),
		server.WithServiceAddr(addr),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.UserServiceName}),
		server.WithRegistry(cli),
	)
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
