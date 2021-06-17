/**
 * @Author: Anpw
 * @Description: 异常捕获处理
 * @File:  recovery
 * @Version: 1.0.0
 * @Date: 2021/6/15 22:48
 */

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"selfblog/global"
	"selfblog/pkg/app"
	"selfblog/pkg/email"
	"selfblog/pkg/errcode"
	"time"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		UserName: global.EmailSetting.UserName,
		PassWord: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
				err := defailtMailer.SendMail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间： %d", time.Now().Unix()),
					fmt.Sprintf("错误信息： %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}
				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
