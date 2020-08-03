package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置当前 context 的超时时间
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		// 覆盖原请求
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
