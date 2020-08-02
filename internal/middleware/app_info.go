package middleware

import "github.com/gin-gonic/gin"

// 服务信息
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "go-blog")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
