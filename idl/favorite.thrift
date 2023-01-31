include "user.thrift"
include "feed.thrift"

namespace go favorite

struct ActionRequest {
  1:required string token                                                   // 凭证token
  2:required string video_id                                                // 视频id
  3:required string action_type (vt.in = "1", vt.in = "2")                  // 操作类型（1-点赞，2-取消点赞）
}

struct ActionResponse {
  1:required i32 status_code               // 状态值
  2:optional string status_msg             // 状态信息
}

struct FavoriteRequest {
  1:required string user_id                // 用户id
  2:required string token                  // 凭证token
}

struct FavoriteResponse {
  1:required i32 status_code               // 状态值
  2:optional string status_msg             // 状态信息
  3:required feed.Video video_list         // 用户点赞视频列表
}

service RelationService{
  ActionResponse ActionFavorite(1: ActionRequest req)            // 点赞操作
  FavoriteResponse GetFavoriteList(1: FavoriteRequest req)       // 获取关注列表
}