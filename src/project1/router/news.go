package router

import (
	"github.com/kataras/iris"
	"project1/logic/V1"
)

//路由层
func News(r iris.Party){
	r.Get("/news", V1.News)
	r.Get("/count", V1.Count)
	r.Get("/sql", V1.Sql)
	r.Get("/add", V1.Add)
	r.Get("/edit", V1.Edit)
	r.Get("/dele", V1.Dele)
	r.Get("/tran", V1.Tran)
	r.Get("/tran1", V1.Tran1)
}
