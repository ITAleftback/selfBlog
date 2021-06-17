/**
 * @Author: Anpw
 * @Description: 超时控制
 * @File:  context_timeout
 * @Version: 1.0.0
 * @Date: 2021/6/16 21:14
 */

package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

//ContextTimeout
/**
 * @Author: Anpw
 * @Description: 设置当前context 超时时间
 * @param t
 * @return func(c *gin.Context)
 */
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
