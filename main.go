package main

import (
	"fmt"
	"github.com/itbread/lotteryport/configer"
	"github.com/itbread/lotteryport/routers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"log"
	"time"
	stdContext "context"
)

func main() {

	path := "conf/config.json"
	config := configer.NewConfiger(path)
	//datasource.InitDb(config)
	app := iris.New()
	app.Use(logMiddleware())
	iris.RegisterOnInterrupt(func() { //服务中断做相应处理
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx)
	})
	tmpl := iris.Django("./web/views", ".html").
		Reload(true) //是否更新
	app.RegisterView(tmpl)
	app.HandleDir("/statics", "./web/statics") //指定静态文件路径

	app.OnAnyErrorCode(func(ctx iris.Context) { //错误页面设置
		Status := ctx.GetStatusCode()
		ctx.ViewData("Status", Status)
		ctx.View("shared/error.html")
	})
	routers.Route(app, config)
	//　使用网络地址启动服务
	portStr := ":80"
	if config.Port > 0 {
		portStr = fmt.Sprintf(":%v", config.Port)
	}
	app.Run(iris.Addr(portStr))
}

// 日志中间件
func logMiddleware() context.Handler {
	return func(ctx iris.Context) {
		r := ctx.Request()
		cookie, _ := r.Cookie("ssq-session")
		log.Printf("Server Request, IP:%s , Cookie:%s , Method: %s , Path: %s ", ctx.RemoteAddr(), cookie, r.Method, r.RequestURI)
		ctx.Next()
	}
}
