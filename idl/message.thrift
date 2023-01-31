namespace go message

struct ChatRequest {
  1:required string token               // 用户鉴权token
  2:required int64 to_user_id           // 对方用户id
}

struct ChatResponse {
  1:required int32 status_code          // 状态码，0-成功，其他值-失败
  2:optional string status_msg          // 返回状态描述
  3:repeated Message message_list       // 消息列表
}

struct Message {
    1:required int64 id                // 消息id
    2:required int64 to_user_id        // 该消息接收者的id
    3:required int64 from_user_id      // 该消息发送者的id
    4:required string content          // 消息内容
    5:optional string create_time      // 消息创建时间
}

struct ActionRequest {
    1:required string token            // 用户鉴权token
    2:required int64 to_user_id        // 对方用户id
    3:required int32 action_type       // 1-发送消息
    4:required string content          // 消息内容
}

struct ActionResponse {
    1:required int32 status_code       // 状态码，0-成功，其他值-失败
    2:optional string status_msg       // 返回状态描述
}

service MessageService{
    ChatResponse ChatMessage(1: ChatRequest req)               // 消息记录
    ActionResponse ActionMessage(1: ActionRequest req)         // 发送消息
}