package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"goblog/global"
	"goblog/pkg/logger"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 实现双写，将副本存储在 body 中
func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		// 替换原有的写入流，扩展一个 body 出来
		c.Writer = bodyWriter
		beginTime := time.Now().Unix()
		// 响应前
		c.Next()
		// 响应后
		endTime := time.Now().Unix()
		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(), // 请求参数
			"response": bodyWriter.body.String(),    // 请求结果
		}
		global.Logger.WithFields(fields).
			Infof("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
				c.Request.Method, bodyWriter.Status(), beginTime, endTime)
		// 控制台打印
		fmt.Printf("access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method, bodyWriter.Status(), beginTime, endTime)
	}
}
