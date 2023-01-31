package service

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/errno"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GetFriendListService struct {
	ctx context.Context
}

func NewGetFriendListService(ctx context.Context) *GetFriendListService {
	return &GetFriendListService{ctx: ctx}
}

func (getFriendListService *GetFriendListService) GetFriendList(ctx context.Context, session neo4j.SessionWithContext, user *user.CheckUserRequest) (int64, error) {
	userid, err := neo4j.ExecuteRead[int64](ctx, session, func(tx neo4j.ManagedTransaction) (int64, error) {
		result, err := tx.Run(ctx, "MATCH (n:User {username: $username}) RETURN n.password AS ps, id(n) AS i LIMIT 1", map[string]any{
			"username": user.Username,
		})
		if err != nil {
			return 0, err
		}
		record, err := result.Single(ctx)
		if err != nil {
			return 0, errno.AuthorizationFailedErr
		}
		return record.Values[1].(int64), nil
	})
	if err != nil {
		return 0, err
	}
	return userid, nil
}
