/**
 * @Author: Anpw
 * @Description:
 * @File:  limiter
 * @Version: 1.0.0
 * @Date: 2021/6/16 20:43
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"selfblog/pkg/app"
	"selfblog/pkg/errcode"
	"selfblog/pkg/limiter"
)

func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
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
