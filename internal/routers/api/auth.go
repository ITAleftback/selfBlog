/**
 * @Author: Anpw
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2021/5/28 22:22
 */

package api

import (
	"github.com/gin-gonic/gin"
	"selfblog/global"
	"selfblog/internal/service"
	"selfblog/pkg/app"
	"selfblog/pkg/errcode"

)

func GetAuth(c *gin.Context)  {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.Errorf(c, "app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf(c, "svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c, "svc.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}


	response.ToResponse(gin.H{"token":token})

}
 
