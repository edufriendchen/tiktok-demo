namespace go user

struct User {
  1:required i64 id                        // 用户id
  2:required string name                   // 用户名称
  3:optional i64 follow_count              // 关注总数
  4:optional i64 follower_count            // 粉丝总数
  5:required bool is_follow                // true-已关注，false-未关注
}

struct CreateUserRequest {
  1:required string username (vt.min_size = "1", vt.max_size = "32")           // 注册用户名，最长32个字符
  2:required string password (vt.min_size = "1", vt.max_size = "32")           // 密码，最长32个字符
}

struct CreateUserResponse {
  1:required i32 status_code               // 状态值
  2:optional string status_msg             // 状态信息
  3:required i64 user_id                   // 用户id
  4:required string token                  // 用户鉴权token
}

struct CheckUserRequest {
  1:string username (vt.min_size = "1", vt.max_size = "32")           // 登录用户名，最长32个字符
  2:string password (vt.min_size = "1", vt.max_size = "32")           // 密码，最长32个字符
}

struct CheckUserResponse {
  1:required i32 status_code               // 状态值
  2:optional string status_msg             // 状态信息
  3:required i64 user_id                   // 用户id
  4:required string token                  // 用户鉴权token
}

struct GetUserRequest {
  1:required i64 user_id                   // 用户id
  2:required string token                  // 用户鉴权token
}

struct GetUserResponse {
  1:required i32 status_code               // 状态值
  2:optional string status_msg             // 状态信息
  3:required User user                     // 用户信息
}

service UserService{
  CreateUserResponse CreateUser(1: CreateUserRequest req)   // 创建用户信息
  CheckUserResponse CheckUser(1: CheckUserRequest req)      // 验证用户
  GetUserResponse GetUser(1: GetUserRequest req)            // 获取用户信息
}   