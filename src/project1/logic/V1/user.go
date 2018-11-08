package V1


import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/imroc/req"
	"github.com/kataras/iris"
	"io/ioutil"
	"project1/conf"
	"project1/model"
)

//数据返回格式
type http_res struct {
	Code int `json:"code"`
	Msg  string	`json:"msg"`
	Data interface{} `json:"data"`
}

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


func Rget(ctx iris.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := client.Set("go_tt", "hello,world!", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("go_tt").Result()
	ctx.JSON(iris.Map{
		"result": val,
		"err": err,
	})
}

func Ask(ctx iris.Context) {
	header := req.Header{
		"Accept":        "application/json",
		//"Authorization": "Basic YWRtaW46YWRtaW4=",
	}
	param := req.Param{
		"code": "023ghu4m1UhEkk0OTX2m1vjd4m1ghu4r",
	}
	// only url is required, others are optional.
	r, _ := req.Post("https://api.sqydt.easysq.cn/api/loginByCode", header, param)
	//ctx.WriteString(r.String())		//直接输出字符串
	res := r.Response()
	body,_ := ioutil.ReadAll(res.Body)
	//ctx.Write(body)

	//结构化请求结果
	xcc := &http_res{}
	json.Unmarshal(body,xcc)
	print(xcc.Code)
	ctx.JSON(xcc)

}
