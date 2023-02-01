// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	relation "github.com/edufriendchen/tiktok-demo/kitex_gen/relation"
)

func serviceInfo() *kitex.ServiceInfo {
	return relationServiceServiceInfo
}

var relationServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "RelationService"
	handlerType := (*relation.RelationService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ActionRelation":   kitex.NewMethodInfo(actionRelationHandler, newRelationServiceActionRelationArgs, newRelationServiceActionRelationResult, false),
		"MGetFollowList":   kitex.NewMethodInfo(mGetFollowListHandler, newRelationServiceMGetFollowListArgs, newRelationServiceMGetFollowListResult, false),
		"MGetFollowerList": kitex.NewMethodInfo(mGetFollowerListHandler, newRelationServiceMGetFollowerListArgs, newRelationServiceMGetFollowerListResult, false),
		"MGetFriendList":   kitex.NewMethodInfo(mGetFriendListHandler, newRelationServiceMGetFriendListArgs, newRelationServiceMGetFriendListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "relation",
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

func actionRelationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceActionRelationArgs)
	realResult := result.(*relation.RelationServiceActionRelationResult)
	success, err := handler.(relation.RelationService).ActionRelation(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceActionRelationArgs() interface{} {
	return relation.NewRelationServiceActionRelationArgs()
}

func newRelationServiceActionRelationResult() interface{} {
	return relation.NewRelationServiceActionRelationResult()
}

func mGetFollowListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceMGetFollowListArgs)
	realResult := result.(*relation.RelationServiceMGetFollowListResult)
	success, err := handler.(relation.RelationService).MGetFollowList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceMGetFollowListArgs() interface{} {
	return relation.NewRelationServiceMGetFollowListArgs()
}

func newRelationServiceMGetFollowListResult() interface{} {
	return relation.NewRelationServiceMGetFollowListResult()
}

func mGetFollowerListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceMGetFollowerListArgs)
	realResult := result.(*relation.RelationServiceMGetFollowerListResult)
	success, err := handler.(relation.RelationService).MGetFollowerList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceMGetFollowerListArgs() interface{} {
	return relation.NewRelationServiceMGetFollowerListArgs()
}

func newRelationServiceMGetFollowerListResult() interface{} {
	return relation.NewRelationServiceMGetFollowerListResult()
}

func mGetFriendListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*relation.RelationServiceMGetFriendListArgs)
	realResult := result.(*relation.RelationServiceMGetFriendListResult)
	success, err := handler.(relation.RelationService).MGetFriendList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newRelationServiceMGetFriendListArgs() interface{} {
	return relation.NewRelationServiceMGetFriendListArgs()
}

func newRelationServiceMGetFriendListResult() interface{} {
	return relation.NewRelationServiceMGetFriendListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ActionRelation(ctx context.Context, req *relation.ActionRequest) (r *relation.ActionResponse, err error) {
	var _args relation.RelationServiceActionRelationArgs
	_args.Req = req
	var _result relation.RelationServiceActionRelationResult
	if err = p.c.Call(ctx, "ActionRelation", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetFollowList(ctx context.Context, req *relation.FollowRequest) (r *relation.FollowResponse, err error) {
	var _args relation.RelationServiceMGetFollowListArgs
	_args.Req = req
	var _result relation.RelationServiceMGetFollowListResult
	if err = p.c.Call(ctx, "MGetFollowList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetFollowerList(ctx context.Context, req *relation.FollowerRequest) (r *relation.FollowerResponse, err error) {
	var _args relation.RelationServiceMGetFollowerListArgs
	_args.Req = req
	var _result relation.RelationServiceMGetFollowerListResult
	if err = p.c.Call(ctx, "MGetFollowerList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGetFriendList(ctx context.Context, req *relation.FriendRequest) (r *relation.FriendResponse, err error) {
	var _args relation.RelationServiceMGetFriendListArgs
	_args.Req = req
	var _result relation.RelationServiceMGetFriendListResult
	if err = p.c.Call(ctx, "MGetFriendList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
