package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// SetResponse pack response
func SetResponse(c *app.RequestContext, response interface{}) {
	c.JSON(consts.StatusOK, response)
}
