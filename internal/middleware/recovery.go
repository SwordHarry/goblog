package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/pkg/app"
	"goblog/pkg/errcode"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Panicf("panic recover err: %v", err)
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
