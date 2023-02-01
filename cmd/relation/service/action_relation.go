package service

import (
	"context"
	"fmt"

	"github.com/edufriendchen/tiktok-demo/kitex_gen/relation"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type ActionRelationService struct {
	ctx     context.Context
	session neo4j.SessionWithContext
}

func NewActionRelationService(ctx context.Context, driver neo4j.DriverWithContext) *ActionRelationService {
	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	return &ActionRelationService{ctx: ctx, session: session}
}

func (s *ActionRelationService) ActionRelation(req *relation.ActionRequest) (int64, error) {
	userid, err := neo4j.ExecuteRead[int64](s.ctx, s.session, func(tx neo4j.ManagedTransaction) (int64, error) {
		result, err := tx.Run(s.ctx, "MATCH (n:User {username: $username}) RETURN n.password AS ps, id(n) AS i LIMIT 1", map[string]any{
			"username": req.ToUserId,
		})
		if err != nil {
			return 0, err
		}
		record, err := result.Single(s.ctx)
		if err != nil {
			return 0, errno.AuthorizationFailedErr
		}
		fmt.Println(record.Values[0], record.Values[1])
		return record.Values[1].(int64), nil
	})
	if err != nil {
		return 0, err
	}
	return userid, nil
}
