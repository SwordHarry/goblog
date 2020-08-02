package api

import (
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	panic("测试邮件发送功能")
	//param := service.AuthRequest{}
	//response := app.NewResponse(c)
	//valid, errs := app.BindAndValid(c, &param)
	//if !valid {
	//	InvalidForBind(response, errs)
	//	return
	//}
	//svc := service.New(c.Request.Context())
	//err := svc.CheckAuth(&param)
	//if err != nil {
	//	global.Logger.Errorf("svc.CheckAuth err: %v", err)
	//	response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
	//	return
	//}
	//token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	//if err != nil {
	//	global.Logger.Errorf("app.GenerateToken err: %v", err)
	//	response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
	//	return
	//}
	//response.ToResponse(gin.H{
	//	"code":  0,
	//	"msg":   "success",
	//	"token": token,
	//})
}
