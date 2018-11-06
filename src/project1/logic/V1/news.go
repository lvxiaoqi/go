package V1

import (
	"fmt"
	"github.com/kataras/iris"
	"project1/model"
)
//查询
func News(ctx iris.Context){
	user := []model.User_info{}
	_map := map[string]interface{}{
		"where" : map[string]interface{}{"t_user.id":1,"t_user_info.user_id":1},
		"field" : []string{"t_user.id","t_user.name","t_user_info.user_loc"},
		"limit"    : map[string]interface{}{
			"limit" : 10,
			"offset" : 0,
		},
		"orderBy" : "t_user.id desc",
		"join" : "LEFT JOIN t_user_info ON t_user_info.user_id = t_user.id",
	}
	res := model.Select(_map,&user)
	ctx.JSON(iris.Map{
		"data": res,
	})
	ctx.HTML("im news list.")
}
//总数
func Count(ctx iris.Context){
	var out int
	user := []model.User_info{}
	_map := map[string]interface{}{
		"where" : map[string]interface{}{},
		"field" : []string{"t_user.id","t_user.name","t_user_info.user_loc"},
		"limit"    : map[string]interface{}{
			"limit" : 10,
			"offset" : 0,
		},
		"orderBy" : "t_user.id desc",
		"join" : "LEFT JOIN t_user_info ON t_user_info.user_id = t_user.id",
	}
	model.Count(_map,&user,&out)
	fmt.Printf("%d",out)
	//ctx.JSON(iris.StatusOK, out)
	ctx.HTML("im total news.")
}
//原生
func Sql(ctx iris.Context){
	type user struct {
		Name string
		ID   int
	}
	res := []user{}
	sql := "select id from t_user where id =1"
	model.Exec(sql,&res)
	fmt.Printf("res:",res)
	ctx.JSON(iris.Map{
		"data": res,
	})
	ctx.HTML("im sql exec.")
}
//新增
func Add(ctx iris.Context){
	user := model.User{
		Name:"galeone9",
	}
	res := model.Add(&user)
	if res.Error != nil {
		fmt.Printf("res:",res.Error)
	}
	ctx.JSON(iris.Map{
		"data": res,
	})
	ctx.HTML("im add one.")
}
//修改
func Edit(ctx iris.Context){
	user := model.User{}
	arr := map[string]interface{}{
		"where":map[string]interface{}{
			"id":6,
		},
		"data":map[string]interface{}{
			"name":"galeone610",
		},
	}
	res := model.Edit(arr,&user)
	if res.Error == nil {
		fmt.Printf("修改成功")
	}else{
		fmt.Printf("res:",res.Error)
	}
	ctx.JSON(iris.Map{
		"data": res,
	})
	ctx.HTML("im edit one.")
}
//删除
func Dele(ctx iris.Context){
	user := model.User{}
	arr := map[string]interface{}{
		"id":8,
	}
	res := model.Dele(arr,user)
	if res.Error == nil {
		fmt.Printf("删除成功")
	}else{
		fmt.Printf("res:",res.Error)
	}
	ctx.JSON(iris.Map{
		"data": res,
	})
	ctx.HTML("im dele one.")
}
//事务
func Tran(ctx iris.Context){
	boo := true
	tx := model.Conn()
	tx = tx.Begin()

	//新增
	user := model.User{
		Name:"galeone14",
	}
	res := model.Tadd(&user,tx)
	if res.Error == nil {
		fmt.Printf("新增")
	}else{
		boo = false
		fmt.Printf("res:",res.Error)
	}
	//修改
	user_info := model.UserInfo{}
	arr := map[string]interface{}{
		"where":map[string]interface{}{
			"user_info_id":3,
		},
		"data":map[string]interface{}{
			"user_loc":"user_loc_地址3",
		},
	}
	res1 := model.Tedit(arr,&user_info,tx)
	if res1.Error == nil {
		fmt.Printf("修改")
	}else{
		boo = false
		fmt.Printf("res:",res1.Error)
	}
	if boo {
		fmt.Printf("this commit")
		tx.Commit()
	}else{
		fmt.Printf("this roolback")
		tx.Rollback()
	}
	ctx.HTML("im transaction.")
}

//事务测试 未封装的方法
func Tran1(ctx iris.Context) {
	boo := true
	tx := model.Conn()
	tx = tx.Begin()
	if err := tx.Create(&model.User{Name: "Giraffe"}).Error; err != nil {
		boo = false
	}

	if err := tx.Create(&model.User{Name: "Lion"}).Error; err != nil {
		boo = false
	}

	if boo {
		fmt.Printf("this commit")
		tx.Commit()
	}else{
		fmt.Printf("this roolback")
		tx.Rollback()
	}

	ctx.HTML("im transaction test.")
}