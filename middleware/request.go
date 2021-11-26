package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xhyonline/xutil/helper"
)

func Static() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.RequestURI()
		if isStatic(path) {
			c.Header("Content-Type", "text/html; charset=utf-8")
		}
		// 处理请求
		c.Next()
	}
}

// isStatic 是否为静态文件
func isStatic(path string) bool {
	return helper.InArray(path, []string{
		"/index", // 首页
	})
}
