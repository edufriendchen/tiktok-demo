package initialize

import (
	"github.com/edufriendchen/tiktok-demo/cmd/user/global"
	"github.com/edufriendchen/tiktok-demo/pkg/consts"
	"github.com/edufriendchen/tiktok-demo/pkg/jwt"
)

// InitJWT to init JWT
func InitJWT() {
	global.Jwt = jwt.NewJWT([]byte(consts.JWTSecretKey))
}
