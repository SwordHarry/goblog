package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/pkg/app"
	"goblog/pkg/email"
	"goblog/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		PassWord: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)
				// 发送崩溃邮件
				err := defailtMailer.SendMail(
					fmt.Sprintf("go-blog 发生异常"),
					fmt.Sprintf("异常抛出，发生时间: %d；错误信息: %v", time.Now().Unix(), err),
					global.EmailSetting.To,
				)
				if err != nil {
					global.Logger.Panicf("panic recover err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
