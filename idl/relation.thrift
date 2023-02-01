include "user.thrift"

namespace go relation

struct ActionRequest {
  1:required string token (vt.min_size = "1")                        // 凭证token
  2:required string to_user_id (vt.min_size = "1")                   // 作用对象用户id
  3:required string action_type (vt.in = "1", vt.in = "2")           // 类型
}

struct ActionResponse {
  1:required i32 status_code                 // 状态值
  2:optional string status_msg               // 状态信息
}

struct FollowRequest {
  1:required i64 user_id                  // 用户id
  2:required string token                    // 凭证token
}

struct FollowResponse {
  1:required i32 status_code                 // 状态值
  2:optional string status_msg               // 状态信息
  3:required list<user.User> user_list       // 关注列表
}

struct FollowerRequest {
  1:required i64 user_id                     // 用户id
  2:required string token                    // 用户鉴权token
}

struct FollowerResponse {
  1:required i32 status_code                 // 状态值
  2:optional string status_msg               // 状态信息
  3:required list<user.User> user_list       // 关注列表
}

struct FriendRequest {
  1:required i64 user_id                     // 用户id
  2:required string token                    // 用户鉴权token
}

struct FriendResponse {
  1:required i32 status_code                 // 状态值
  2:optional string status_msg               // 状态信息
  3:required list<user.User> user_list       // 粉丝列表
}

service RelationService{
  ActionResponse ActionRelation(1: ActionRequest req)          // 关注操作
  FollowResponse MGetFollowList(1: FollowRequest req)          // 获取关注列表
  FollowerResponse MGetFollowerList(1: FollowerRequest req)    // 获取粉丝列表
  FriendResponse MGetFriendList(1: FriendRequest req)          // 获取好友列表
}