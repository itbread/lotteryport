package handlers

import "github.com/kataras/iris/v12"

func (h *homeHandler) Home(ctx iris.Context) {
	ctx.View("pages/login/login.html", iris.Map{
		"curPage":     "系统首页",
		"curUrl":      ctx.Path(),
		"Title":  "首页",
		"secondTitle": "系统首页",
	})
}
