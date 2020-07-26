package v1

import (
	"goblog/global"
	"goblog/pkg/app"
	"goblog/pkg/errcode"
)

// 封装请求中的重复代码
// 写入日志，并将错误信息写入 response
func invalidForBind(response *app.Response, errs app.ValidErrors) {
	global.Logger.Errorf("app.BindAndValid errs: %v", errs)
	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
}
