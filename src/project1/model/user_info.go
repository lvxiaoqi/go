package model

import (
	"github.com/jinzhu/gorm"
)

type UserInfo struct {
	Id   int   `gorm:"column:user_info_id"`
	Loc string `gorm:"column:user_loc"`	//数据库中的真实的字段名
}
//初始化
func init(){
	var db *gorm.DB = Conn()
	db.SingularTable(false)
}

//表名
func (UserInfo) TableName() string {
	return "t_user_info"
}

