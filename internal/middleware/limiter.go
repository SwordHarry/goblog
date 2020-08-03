package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog/pkg/app"
	"goblog/pkg/errcode"
	"goblog/pkg/limiter"
)

// 限流中间件
func RateLimiter(l limiter.LimitInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// 获取 1 个令牌
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
