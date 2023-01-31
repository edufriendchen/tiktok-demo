package service

import (
	"context"
	"fmt"

	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (getUserService *GetUserService) GetUser(ctx context.Context, session neo4j.SessionWithContext, param *user.GetUserRequest) (*user.User, error) {
	user, err := neo4j.ExecuteRead[*user.User](ctx, session, func(tx neo4j.ManagedTransaction) (*user.User, error) {
		result, err := tx.Run(ctx, "MATCH (n:User) WHERE id(n) = $toUserId MATCH (n1:User) WHERE id(n1) = $userId RETURN n, CASE WHEN (n1)-[:follow]->(n) THEN true ELSE false END AS result",
			map[string]any{
				"toUserId": 0,
				"userId":   param.UserId,
			})
		if err != nil {
			return nil, err
		}
		record, err := result.Single(ctx)
		if err != nil {
			return nil, err
		}
		rawItemNode, found := record.Get("n")
		if !found {
			return nil, fmt.Errorf("could not find column")
		}
		itemNode := rawItemNode.(neo4j.Node)
		id := itemNode.GetId()
		if err != nil {
			return nil, err
		}
		name, err := neo4j.GetProperty[string](itemNode, "nickname")
		if err != nil {
			return nil, err
		}
		follow_count, err := neo4j.GetProperty[int64](itemNode, "followCount")
		if err != nil {
			return nil, err
		}
		follower_count, err := neo4j.GetProperty[int64](itemNode, "followerCount")
		if err != nil {
			return nil, err
		}
		IsFollow, found := record.Values[1].(bool)
		if err != nil {
			return nil, err
		}
		return &user.User{Id: id, Name: name, FollowCount: &follow_count, FollowerCount: &follower_count, IsFollow: IsFollow}, nil
	})
	return user, err
}
