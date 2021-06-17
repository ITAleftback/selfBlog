/**
 * @Author: Anpw
 * @Description:
 * @File:  dao
 * @Version: 1.0.0
 * @Date: 2021/5/26 22:24
 */

package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(engine *gorm.DB) *Dao {
	return &Dao{engine:engine}
}
 
