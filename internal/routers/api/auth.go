package api

import (
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/internal/request"
	"goblog/internal/routers/common"
	"goblog/internal/service"
	"goblog/pkg/app"
	"goblog/pkg/errcode"
)

func GetAuth(c *gin.Context) {
	param := request.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		common.InvalidForBind(c, response, errs)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"code":  0,
		"msg":   "success",
		"token": token,
	})
}
