package initialize

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/cmd/user/global"

	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func InitNeo4j() error {
	driver, err := neo4j.NewDriverWithContext(consts.Neo4jUri, neo4j.BasicAuth("neo4j", "friendchen0429", ""))
	if err != nil {
		return err
	}
	ctx := context.Background()
	defer driver.Close(ctx)
	global.Neo4jSession = driver.NewSession(ctx, neo4j.SessionConfig{})
	defer global.Neo4jSession.Close(ctx)
	return nil
}
