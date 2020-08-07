package main

import (
	"github.com/itbread/lotteryport/configer"
	"github.com/itbread/lotteryport/datasource"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/itbread/lotteryport/controllers"
)

func main() {
	path := "conf/config.json"
	config := configer.NewConfiger(path)
	datasource.InitPgsqlDb(config)
	app := iris.New()
	tmpl := iris.Django("./templates/views", ".html").Reload(true) //是否更新
	app.RegisterView(tmpl)
	app.HandleDir("/statics", "./templates/statics") //指定静态文件路径
	users := mvc.New(app.Party("/ssq"))              // 一个根路径为 /users 的组
	booksAPI := app.Party("/books")
	{
		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}
	users.Handle(new(controllers.SsqController)) // 定义组的controller
	app.Get("/", index)
	app.Post("/data", GetdataList)
	//　使用网络地址启动服务
	app.Run(iris.Addr(":9090"))
}

func index(ctx iris.Context) {

}

type Book struct {
	Title string `json:"title"`
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
func GetdataList(ctx iris.Context) {

}
