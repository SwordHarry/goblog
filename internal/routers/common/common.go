package common

import (
	"context"
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/pkg/app"
	"goblog/pkg/errcode"
	"net/http"
)

// 封装请求中的重复代码
// 写入日志，并将错误信息写入 response
func InvalidForBind(ctx context.Context, response *app.Response, errs app.ValidErrors) {
	global.Logger.Errorf(ctx, "app.BindAndValid errs: %v", errs)
	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
}

// 成功
func View(c *gin.Context, html string, data gin.H) {
	c.HTML(http.StatusOK, html, data)
}

// 错误
func ViewErr(c *gin.Context, err error) {
	c.HTML(http.StatusInternalServerError, "error.html", gin.H{
		"error": err,
	})
}
