package model

import (
	"github.com/jinzhu/gorm"
)

//数据库结构
type User struct {
	//Model
	Name string
	ID   int
}
type User_info struct {
	//Model	用于联查
	Name string
	ID   int
	Loc string `gorm:"column:user_loc"`	//数据库中的真实的字段名
}

//初始化
func init(){
	var db *gorm.DB = Conn()
	db.SingularTable(false)
}

//表名
func (User) TableName() string {
	return "t_user"
}
//表名
func (User_info) TableName() string {
	return "t_user"
}

