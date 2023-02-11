package main

import (
	"context"
	comment "github.com/edufriendchen/tiktok-demo/kitex_gen/comment"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// ActionComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) ActionComment(ctx context.Context, req *comment.ActionRequest) (resp *comment.ActionResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetCommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) MGetCommentList(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	// TODO: Your code here...
	return
}
