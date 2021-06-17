/**
 * @Author: Anpw
 * @Description: 服务信息存储
 * @File:  app_info
 * @Version: 1.0.0
 * @Date: 2021/6/16 14:22
 */

package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}