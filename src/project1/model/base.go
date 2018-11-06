package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"project1/conf"
)
//基础结构
type Model struct {
	ID        uint `gorm:"primary_key"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
	//DeletedAt *time.Time	//当存在该字段时软删除  执行删除操作时将该字段的时间修改为当前时间
}
//链接
func Conn() *gorm.DB{
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"  设置上海时区
	conn := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":3306)/" + conf.DB_DATABASE + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	return db
}

/*
查询多条
arr = map[string]interface {
	"where":map[string]interface {},		//条件
	"or"   :map[string]interface {},		//或查询
	"field":[]string{},				//查询的字段
	"not"  :map[string]interface{}{			//非查询
		"field" : string,			//字段名称
		"value" : string,			//条件
	},
	"in"	: map[string]interface{}{		//in查询
		"field" : "id",
		"value" : []int{1,2,3,4},
	},
	"notin" : map[string]interface{}{		//notin查询
		"field" : "id",
		"value" : []int{4},
	},
	"like"  : map[string]interface{}{		//like查询
		"field" : "name",
		"value" : "Jin%",
	},
	"exp"   : map[string]interface{}{		//条件查询
		"field" :"id",
		"condition" : ">",
		"value" : "1",
	},
	"limit"    : map[string]interface{}{		//限制查询
		"limit" : 10,				//限制条数
		"offset" : 0,				//起始
	},
	"orderBy" : "id desc",				//排序
	"join" : "",					//联查

}
tb 表的结构体
*/
func Select(arr map[string]interface{},tb interface{}) *gorm.DB{
	db := Conn()
	defer db.Close()
	if _,ok := arr["join"]; ok {
		db = db.Joins(arr["join"].(string))
	}
	if _,ok := arr["where"]; ok {
		where := arr["where"].(map[string]interface{})
		for k , v := range where {
			db = db.Where(k + " = ? ",v)
		}
	}
	if _,ok := arr["or"]; ok {
		db = db.Or(arr["or"])
	}
	if _,ok := arr["field"]; ok {
		db = db.Select(arr["field"])
	}
	if _,ok := arr["not"]; ok {
		not:= arr["not"].(map[string]interface{})
		db = db.Where(not["field"].(string) + " <> ?", not["value"])
	}
	if _,ok := arr["in"]; ok {
		in := arr["in"].(map[string]interface{})
		db = db.Where(in["field"].(string) + " IN (?)", in["value"])
	}
	if _,ok := arr["notin"]; ok {
		notin := arr["notin"].(map[string]interface{})
		db = db.Where(notin["field"].(string) + " NOT IN (?)", notin["value"])
	}
	if _,ok := arr["like"]; ok {
		like := arr["like"].(map[string]interface{})
		db = db.Where(like["field"].(string) + " LIKE ?", like["value"])
	}
	if _,ok := arr["exp"]; ok {
		exp := arr["exp"].(map[string]interface{})
		db = db.Where(exp["field"].(string) + " " + exp["condition"].(string) + " " + " ? ", exp["value"])
	}
	if _,ok := arr["limit"]; ok {
		limit := arr["limit"].(map[string]interface{})
		db = db.Limit(limit["limit"])
		db = db.Offset(limit["offset"])
	}
	if _,ok := arr["orderBy"]; ok {
		db = db.Order(arr["orderBy"].(string))
	}
	obj := db.Find(tb)
	return obj
}

/*
查询总数
arr = map[string]interface {
	"where":map[string]interface {},		//条件
	"or"   :map[string]interface {},		//或查询
	"field":[]string{},				//查询的字段
	"not"  :map[string]interface{}{			//非查询
		"field" : string,			//字段名称
		"value" : string,			//条件
	},
	"in"	: map[string]interface{}{		//in查询
		"field" : "id",
		"value" : []int{1,2,3,4},
	},
	"notin" : map[string]interface{}{		//notin查询
		"field" : "id",
		"value" : []int{4},
	},
	"like"  : map[string]interface{}{		//like查询
		"field" : "name",
		"value" : "Jin%",
	},
	"exp"   : map[string]interface{}{		//条件查询
		"field" :"id",
		"condition" : ">",
		"value" : "1",
	},
	"limit"    : map[string]interface{}{		//限制查询
		"limit" : 10,				//限制条数
		"offset" : 0,				//起始
	},
	"orderBy" : "id desc",				//排序
	"join" : "",					//联查

}
tb 表的结构体
res 结果/int
*/
func Count(arr map[string]interface{},tb interface{},res interface{}) *gorm.DB{
	db := Conn()
	defer db.Close()
	if _,ok := arr["join"]; ok {
		db = db.Joins(arr["join"].(string))
	}
	if _,ok := arr["where"]; ok {
		where := arr["where"].(map[string]interface{})
		for k , v := range where {
			db = db.Where(k + " = ? ",v)
		}
	}
	if _,ok := arr["or"]; ok {
		db = db.Or(arr["or"])
	}
	if _,ok := arr["field"]; ok {
		db = db.Select(arr["field"])
	}
	if _,ok := arr["not"]; ok {
		not:= arr["not"].(map[string]interface{})
		db = db.Where(not["field"].(string) + " <> ?", not["value"])
	}
	if _,ok := arr["in"]; ok {
		in := arr["in"].(map[string]interface{})
		db = db.Where(in["field"].(string) + " IN (?)", in["value"])
	}
	if _,ok := arr["notin"]; ok {
		notin := arr["notin"].(map[string]interface{})
		db = db.Where(notin["field"].(string) + " NOT IN (?)", notin["value"])
	}
	if _,ok := arr["like"]; ok {
		like := arr["like"].(map[string]interface{})
		db = db.Where(like["field"].(string) + " LIKE ?", like["value"])
	}
	if _,ok := arr["exp"]; ok {
		exp := arr["exp"].(map[string]interface{})
		db = db.Where(exp["field"].(string) + " " + exp["condition"].(string) + " " + " ? ", exp["value"])
	}
	if _,ok := arr["limit"]; ok {
		limit := arr["limit"].(map[string]interface{})
		db = db.Limit(limit["limit"])
		db = db.Offset(limit["offset"])
	}
	if _,ok := arr["orderBy"]; ok {
		db = db.Order(arr["orderBy"].(string))
	}
	obj := db.Model(tb).Count(res)
	return obj
}

/*
原生执行
type user struct {
		Name string
		ID   int
}
res := []user{}			//输出的格式
sql := "select id from t_user where id =1"	//sql语句
model.Exec(sql,&res)
*/
func Exec(sql string,res interface{}) *gorm.DB{
	db := Conn()
	defer db.Close()
	obj := db.Raw(sql).Scan(res)
	return obj
}
/*
添加
arr := model.User{
		Name:"galeone",
	}
model.Add(&user)
arr 为结构
*/
func Add(arr interface{}) *gorm.DB{
	db := Conn()
	defer db.Close()
	obj := db.Create(arr)
	return obj
}
/*
修改
tb := model.User{}	//数据库
arr := map[string]interface{}{
	"where":map[string]interface{}{		//条件
		"id":6,
	},
	"data":map[string]interface{}{		//修改的数据
		"name":"galeone6",
	},
}
res := model.Edit(arr,&user)
if res.Error == nil {
	fmt.Printf("修改成功")
}
*/
func Edit(arr map[string]interface{},tb interface{}) *gorm.DB{
	db := Conn()
	defer db.Close()
	if _,ok := arr["where"]; ok {
		where := arr["where"].(map[string]interface{})
		for k , v := range where {
			db = db.Where(k + " = ? ",v)
		}
	}
	data := arr["data"]
	obj := db.Model(tb).Updates(data)
	return obj
}

/*
删除
tb := model.User{}	//数据库
arr := map[string]interface{}{	//条件
	"id":7,
}
model.Dele(arr,tb)
*/
func Dele(arr map[string]interface{},tb interface{}) *gorm.DB{
	db := Conn()
	defer db.Close()
	for k , v := range arr {
		db = db.Where(k + " = ? ",v)
	}
	obj := db.Delete(tb)
	return obj
}

/*
用于事务情况下的添加
arr := model.User{
		Name:"galeone",
	}
model.Add(&user)
arr 为结构
db 开始了事务的db对象
*/
func Tadd(arr interface{},tdb *gorm.DB) *gorm.DB{
	obj := tdb.Create(arr)
	return obj
}
/*
用于事务情况下的修改
tb := model.User{}	//数据库
arr := map[string]interface{}{
	"where":map[string]interface{}{		//条件
		"id":6,
	},
	"data":map[string]interface{}{		//修改的数据
		"name":"galeone6",
	},
}
res := model.Edit(arr,&user)
if res.Error == nil {
	fmt.Printf("修改成功")
}
db 开始了事务的db对象
*/
func Tedit(arr map[string]interface{},tb interface{},tdb *gorm.DB) *gorm.DB{
	if _,ok := arr["where"]; ok {
		where := arr["where"].(map[string]interface{})
		for k , v := range where {
			tdb = tdb.Where(k + " = ? ",v)
		}
	}
	data := arr["data"]
	obj := tdb.Model(tb).Updates(data)
	return obj
}

/*
用于事务情况下的删除
tb := model.User{}	//数据库
arr := map[string]interface{}{	//条件
	"id":7,
}
model.Dele(arr,tb)
db 开始了事务的db对象
*/
func Tdele(arr map[string]interface{},tb interface{},tdb *gorm.DB) *gorm.DB{
	for k , v := range arr {
		tdb = tdb.Where(k + " = ? ",v)
	}
	obj := tdb.Delete(tb)
	return obj
}
//过滤

//验证





