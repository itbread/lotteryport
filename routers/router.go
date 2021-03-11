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
	homeParty := app.Party("/")
	homeRoute(homeParty, newHandlers)

	ssqParty := v1Party.Party("/ssq")
	ssqRoute(ssqParty, newHandlers)

	dltParty := v1Party.Party("/dlt")
	dltRoute(dltParty, newHandlers)

	kl8Party := v1Party.Party("/kl8")
	kl8Route(kl8Party, newHandlers)
}

func homeRoute(party iris.Party, newHandlers *handlers.Handlers) {
	party.Get("/login", newHandlers.HomeHandler.Home)
}
func ssqRoute(party iris.Party, newHandlers *handlers.Handlers) {
	party.Get("/history", newHandlers.SsqHandler.GetSsqHistory)
	party.Get("/list", newHandlers.SsqHandler.GetSsqList)
}
func kl8Route(party iris.Party, newHandlers *handlers.Handlers) {
	party.Get("/history", newHandlers.Kl8Handler.GetKl8History)
}
func dltRoute(party iris.Party, newHandlers *handlers.Handlers) {
	party.Get("/history", newHandlers.DltHandler.GetDltHistory)
	party.Get("/list", newHandlers.DltHandler.GetDltList)
	party.Get("/save", newHandlers.DltHandler.SaveDltHistory)
}
