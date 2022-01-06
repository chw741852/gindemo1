package middleware

import (
	"fmt"
	"test/internal/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer logger.Log.Sync() // flushes buffer, if any

		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")
		// 请求前
		c.Next()
		// 请求后
		latency := time.Since(t)
		// 获取发送的 status
		status := c.Writer.Status()
		logger.Log.Info(fmt.Sprintf("%s|%d", latency, status))
	}
}