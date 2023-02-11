package dal

import (
	"context"
	"fmt"

	"github.com/edufriendchen/tiktok-demo/kitex_gen/feed"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/edufriendchen/tiktok-demo/pkg/global"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func MGetPublishListLimit(ctx context.Context, session neo4j.SessionWithContext, req *feed.FeedRequest, limit int64) ([]*feed.Video, error) {
	var cypher = "MATCH (u:User), (v:Video) WHERE (u)-[:publish]->(v) RETURN v,u,false AS k LIMIT $num"
	var params = map[string]any{
		"num": limit,
	}
	claims, err := global.Jwt.ParseToken("fdsfsgsdg")
	if err == nil {
		cypher = "MATCH (u:User), (v:Video) WHERE (u)-[:publish]->(v) RETURN v,u, id(u) = $user_id AS k LIMIT $num"
		params = map[string]any{
			"num":     limit,
			"user_id": claims.Id,
		}
	}
	var videoList []*feed.Video
	var videoItem feed.Video
	var userItem user.User
	fmt.Println("执行语句:", cypher, params)
	_, err = neo4j.ExecuteRead[bool](ctx, session, func(tx neo4j.ManagedTransaction) (bool, error) {
		result, err := tx.Run(ctx, cypher, params)
		if err != nil {
			return false, err
		}
		record, err := result.Collect(ctx)
		if err != nil {
			return false, err
		}
		for i := 0; i < len(record); i++ {
			video_value, ok := record[i].Get("v")
			if ok {
				itemNode := video_value.(neo4j.Node)
				videoItem.Id = itemNode.GetId()
				if videoItem.Title, err = neo4j.GetProperty[string](itemNode, "title"); err != nil {
					return false, err
				}
				if videoItem.CoverUrl, err = neo4j.GetProperty[string](itemNode, "cover_url"); err != nil {
					return false, err
				}
				if videoItem.PlayUrl, err = neo4j.GetProperty[string](itemNode, "play_url"); err != nil {
					return false, err
				}
				if videoItem.CommentCount, err = neo4j.GetProperty[int64](itemNode, "comment_count"); err != nil {
					return false, err
				}
				if videoItem.FavoriteCount, err = neo4j.GetProperty[int64](itemNode, "favorite_count"); err != nil {
					return false, err
				}
			}
			user_value, ok := record[i].Get("u")
			if ok {
				itemNode := user_value.(neo4j.Node)
				userItem.Id = itemNode.GetId()
				if userItem.Name, err = neo4j.GetProperty[string](itemNode, "nickname"); err != nil {
					return false, err
				}
				if userItem.Avatar, err = neo4j.GetProperty[string](itemNode, "avatar"); err != nil {
					return false, err
				}
				follow_count, err := neo4j.GetProperty[int64](itemNode, "follow_count")
				if err != nil {
					return false, err
				}
				follower_count, err := neo4j.GetProperty[int64](itemNode, "follower_count")
				if err != nil {
					return false, err
				}
				userItem.FollowCount = &follow_count
				userItem.FollowerCount = &follower_count
			}
			favorite_value, ok := record[i].Get("k")
			fmt.Println("favorite_value:", favorite_value)
			videoItem.Author = &userItem
			videoList = append(videoList, &videoItem)
		}
		return true, nil
	})
	return videoList, err
}
