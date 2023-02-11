package dal

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/kitex_gen/feed"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/user"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateVideo(ctx context.Context, session neo4j.SessionWithContext, user_id int64, to_user_id int64) (bool, error) {
	is, err := neo4j.ExecuteRead[bool](ctx, session, func(tx neo4j.ManagedTransaction) (bool, error) {
		result, err := tx.Run(ctx, "MATCH (a:User) WHERE id(a) = $user_id MATCH (b:User) WHERE id(b) = $to_user_id MATCH (a)-[f:follow]->(b) WITH COUNT(f) > 0  as is_follow RETURN is_follow", map[string]any{
			"user_id":    user_id,
			"to_user_id": to_user_id,
		})
		if err != nil {
			return false, err
		}
		record, err := result.Single(ctx)
		if err != nil {
			return false, err
		}
		return record.Values[0].(bool), nil
	})
	return is, err
}

func MGetPublishListById(ctx context.Context, session neo4j.SessionWithContext, user_id int64) ([]*feed.Video, error) {
	var videoList []*feed.Video
	var videoItem feed.Video
	var userItem user.User
	_, err := neo4j.ExecuteRead[bool](ctx, session, func(tx neo4j.ManagedTransaction) (bool, error) {
		result, err := tx.Run(ctx, "MATCH (u:User), (v:Video) WHERE id(u) = $user_id AND (u)-[:publish]->(v) RETURN v,u", map[string]any{
			"user_id": user_id,
		})
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
			videoItem.Author = &userItem
			videoList = append(videoList, &videoItem)
		}
		return true, nil
	})
	return videoList, err
}
