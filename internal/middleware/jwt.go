package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goblog/pkg/app"
	"goblog/pkg/errcode"
)

// jwt 中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		// 先从 url query 中寻找token
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			// 从请求头中寻找 token
			token = c.GetHeader("token")
		}

		if token == "" {
			ecode = errcode.UnauthorizedTokenError
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired: // 超时
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			// 中间件中断，返回
			c.Abort()
			return
		}
		// 正常工作，next
		c.Next()
	}
}
