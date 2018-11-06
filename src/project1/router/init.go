package router

import (
	"fmt"
	"github.com/kataras/iris"
)

/*
	路由初始化 并做路由分组
	2017-02-23
*/

func Init(app *iris.Application){
	//版本1
	r := app.Party("/v1")
	//认证
	ra := app.Party("/v1",func(ctx iris.Context){
		fmt.Printf("this auth == ")
		ctx.Next()
	})
	News(r)
	User(ra)

	//版本2
	r2 := app.Party("/v2")
	User2(r2)
}

