include "feed.thrift"

namespace go publish

struct ActionRequest {
    1:required bytes data                                                     // 视频数据
    2:required string token                                                   // 用户鉴权token
    3:required string title (vt.in = "1", vt.in = "2")                        // 视频标题
}

struct ActionResponse {
    1:required i32 status_code               // 状态值
    2:optional string status_msg             // 状态信息
}

struct PublishRequest {
    1:required string user_id                // 用户id
    2:required string token                  // 凭证token
}

struct PublishResponse {
    1:required i32 status_code               // 状态值
    2:optional string status_msg             // 状态信息
    3:required list<feed.Video> video_list   // 用户发布的视频列表
}

service RelationService{
  ActionResponse ActionPulish(1: ActionRequest req)              // 视频上传操作
  CommentResponse GetCommentList(1: CommentRequest req)          // 获取发布的视频列表
}