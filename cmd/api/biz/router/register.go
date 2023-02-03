// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/edufriendchen/tiktok-demo/cmd/api/biz/handler"
	
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!

	root := r.Group("/douyin")

		{
			_user := root.Group("/user")
			_user.GET("/",  handler.MGetUserInfo)
			_user.POST("/login/",  handler.Login)
			_user.POST("/register/", handler.Register)
		}

		{
			_relation := root.Group("/relation")
			_relation.POST("/action",  handler.RelationAction)
			_relation.GET("/follow/list",  handler.MGetFollowList)
			_relation.GET("/follower/list", handler.MGetFollowerList)
			_relation.GET("/friend/list", handler.MGetFriendList)
		}
}
