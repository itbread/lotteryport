package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/itbread/lotteryport/services"
)

type SsqController struct {
	Ctx     iris.Context
	Service services.SsqService
}
