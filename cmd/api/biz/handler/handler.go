package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	c.JSON(consts.StatusOK, utils.H{
		"message": "pong",
	})
}

// SetResponse pack response
func SetResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}
