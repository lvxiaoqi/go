package router

import (
	"github.com/kataras/iris"
	"project1/logic/V1"
	"project1/logic/V2"
)
//第一版本
func User(r iris.Party){
	r.Get("/user", V1.List)
	r.Get("/", V1.Index)
	r.Get("/tt", V1.Tt)
	r.Get("/myjson", V1.Myjson)
	r.Get("/redis", V1.Rget)
	r.Get("/ask", V1.Ask)
}
//第二版本
func User2(r iris.Party){
	r.Get("/user", V2.List)
}

