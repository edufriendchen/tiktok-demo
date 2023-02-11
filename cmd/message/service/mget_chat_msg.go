package service

import (
	"context"

	"github.com/edufriendchen/tiktok-demo/cmd/message/dal"
	"github.com/edufriendchen/tiktok-demo/cmd/message/pack"
	"github.com/edufriendchen/tiktok-demo/kitex_gen/message"
)

type MGetChatMsgService struct {
	ctx context.Context
}

func NewMGetChatMsgService(ctx context.Context) *MGetChatMsgService {
	return &MGetChatMsgService{ctx: ctx}
}

func (s *MGetChatMsgService) MGetChatMsg(req *message.ChatRequest, to_user_id int64) ([]*message.Message, error) {
	modelMessages, err := dal.MGetChatMsg(s.ctx, to_user_id, req.ToUserId)
	if err != nil {
		return nil, err
	}
	return pack.Messages(modelMessages), nil
}
