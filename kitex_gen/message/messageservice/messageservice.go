// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	message "github.com/edufriendchen/tiktok-demo/kitex_gen/message"
)

func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

var messageServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MessageService"
	handlerType := (*message.MessageService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ChatMessage":   kitex.NewMethodInfo(chatMessageHandler, newMessageServiceChatMessageArgs, newMessageServiceChatMessageResult, false),
		"ActionMessage": kitex.NewMethodInfo(actionMessageHandler, newMessageServiceActionMessageArgs, newMessageServiceActionMessageResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "message",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func chatMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*message.MessageServiceChatMessageArgs)
	realResult := result.(*message.MessageServiceChatMessageResult)
	success, err := handler.(message.MessageService).ChatMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceChatMessageArgs() interface{} {
	return message.NewMessageServiceChatMessageArgs()
}

func newMessageServiceChatMessageResult() interface{} {
	return message.NewMessageServiceChatMessageResult()
}

func actionMessageHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*message.MessageServiceActionMessageArgs)
	realResult := result.(*message.MessageServiceActionMessageResult)
	success, err := handler.(message.MessageService).ActionMessage(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageServiceActionMessageArgs() interface{} {
	return message.NewMessageServiceActionMessageArgs()
}

func newMessageServiceActionMessageResult() interface{} {
	return message.NewMessageServiceActionMessageResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ChatMessage(ctx context.Context, req *message.ChatRequest) (r *message.ChatResponse, err error) {
	var _args message.MessageServiceChatMessageArgs
	_args.Req = req
	var _result message.MessageServiceChatMessageResult
	if err = p.c.Call(ctx, "ChatMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ActionMessage(ctx context.Context, req *message.ActionRequest) (r *message.ActionResponse, err error) {
	var _args message.MessageServiceActionMessageArgs
	_args.Req = req
	var _result message.MessageServiceActionMessageResult
	if err = p.c.Call(ctx, "ActionMessage", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
