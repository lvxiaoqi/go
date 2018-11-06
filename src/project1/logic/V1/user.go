package V1


import (
	"fmt"
	"github.com/kataras/iris"
	"project1/conf"
	"project1/model"
)


func List(ctx iris.Context) {
	//db := model.Conn()
	//db.Close()
	//var user model.User
	user := []model.User{}
	/*type Result struct {
		Name string
		ID  int
	}
	var result Result*/
	//res := db.Find(&user)
	/*var _map map[string]interface{}
	_map = make(map[string]interface{})
	_map["where"] = map[string]interface{}{
		"id":3,
	}*/
	_map := map[string]interface{}{
		"where" : map[string]interface{}{"id":1},
	}
	res := model.Select(_map,user)
	ctx.JSON(iris.Map{
		"data": res,
	})

	fmt.Println("result:", res.Rows)
	//model.One()
	//fmt.Printf(res.Value)
	ctx.HTML("user list v1.")
}

func Index(ctx iris.Context) {
	ctx.HTML("hello World!")
}

func Tt(ctx iris.Context) {
	ctx.HTML("hello World!")
	fmt.Printf("hello, world\n")
}

func Myjson(ctx iris.Context) {
	wid := conf.WIDTH
	len := conf.LENGTH
	s := wid*len
	ctx.JSON(iris.Map{
		"Name": "Iris",
		"Released": "13 March 2016",
		"Stars": s,
	})
}

