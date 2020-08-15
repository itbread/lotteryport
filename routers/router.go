package routers

import (
	"github.com/itbread/lotteryport/configer"
	"github.com/itbread/lotteryport/handlers"
	"github.com/kataras/iris/v12"
)

func Route(app *iris.Application, config *configer.Config) {
	v1Party := app.Party("/v1")
	newHandlers := handlers.NewHandler(config)
	// 主页相关路由
	ssqParty := v1Party.Party("/ssq")
	homeRoute(ssqParty, newHandlers)
	dltParty := v1Party.Party("/dlt")
	dltRoute(dltParty, newHandlers)
}

func homeRoute(party iris.Party, newHandlers *handlers.Handlers) {
	party.Get("/history", newHandlers.SsqHandler.GetSsqHistory)
	party.Get("/list", newHandlers.SsqHandler.GetSsqList)
}
func dltRoute(party iris.Party, newHandlers *handlers.Handlers) {
	party.Get("/history", newHandlers.DltHandler.GetDltHistory)
	party.Get("/list", newHandlers.DltHandler.GetDltList)
	party.Get("/save", newHandlers.DltHandler.SaveDltHistory)
}
