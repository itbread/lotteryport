package main

import (
	"encoding/json"
	"fmt"
	"github.com/itbread/lotteryport/datamodels"
	"github.com/itbread/lotteryport/utils/http"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/itbread/lotteryport/controllers"
)


func main() {
	app := iris.New()
	tmpl := iris.Django("./templates/views", ".html").Reload(true) //是否更新
	app.RegisterView(tmpl)
	app.StaticWeb("/statics", "./templates/statics") //指定静态文件路径
	users := mvc.New(app.Party("/ssq"))  // 一个根路径为 /users 的组

	users.Handle(new(controllers.SsqController)) // 定义组的controller
var ssq datamodels.Ssq
	ssq.Code="20190616"
	bytes,err:=json.Marshal(ssq)
	if err!=nil {
		
	}
	fmt.Println(string(bytes))
	http.Get()
	app.Get("/", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		ctx.ViewData("message", "Hello world!")
		// 渲染模板文件：
		ctx.View("index.html")
	})
	//　使用网络地址启动服务
	app.Run(iris.Addr(":9090"))
}