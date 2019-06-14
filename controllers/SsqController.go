package controllers

import (
	"github.com/kataras/iris"
	"lotteryweb/services"
)

type SsqController struct {
	Ctx     iris.Context
	Service services.SsqService
}
