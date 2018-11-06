package main

import (
	"github.com/kataras/iris"
	"project1/router"
)
func main(){
	app := iris.Default()
	router.Init(app)
	app.Run(iris.Addr(":9091"))
}

