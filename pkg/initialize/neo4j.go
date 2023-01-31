package initialize

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/edufriendchen/tiktok-demo/cmd/user/global"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func InitNeo4j() {
	driver, err := neo4j.NewDriverWithContext(consts.Neo4jUri, neo4j.BasicAuth("neo4j", "friendchen0429", ""))
	if err != nil {
		hlog.Fatalf("init neo4j failed: %s", err.Error())
	}
	ctx := context.Background()
	global.Neo4jSession = driver.NewSession(ctx, neo4j.SessionConfig{})
}
