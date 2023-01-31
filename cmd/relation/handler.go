package main

import (
	"context"

	relation "github.com/edufriendchen/tiktok-demo/kitex_gen/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// ActionRelation implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) ActionRelation(ctx context.Context, req *relation.ActionRequest) (resp *relation.ActionResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.FollowRequest) (resp *relation.FollowResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.FollowerRequest) (resp *relation.FollowerResponse, err error) {
	// TODO: Your code here...
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.FriendRequest) (resp *relation.FriendResponse, err error) {
	// TODO: Your code here...
	return
}
