/**
 * @Author: Anpw
 * @Description:
 * @File:  service
 * @Version: 1.0.0
 * @Date: 2021/5/26 22:44
 */

package service

import (
	"context"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"selfblog/global"
	"selfblog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx:ctx}
	svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))
	return svc
}



 
