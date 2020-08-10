package handlers

import (
	"github.com/itbread/lotteryport/configer"
	"github.com/itbread/lotteryport/datasource"
	"github.com/itbread/lotteryport/services"
	"github.com/jinzhu/gorm"
	"log"
)

type Handler struct {
	config *configer.Config
	db     *gorm.DB
}

type ssqHandler struct {
	Handler
	service services.SsqService
}

type Handlers struct {
	SsqHandler *ssqHandler
}

func NewHandler(config *configer.Config) *Handlers {
	db, err := datasource.InitDb(config)
	if err != nil {
		log.Panicf("Error init data orm: %v", err)
		db.Close()
		return nil
	}
	ssqService := services.NewSsqService(db)
	basehand := Handler{config: config, db: db}
	return &Handlers{
		SsqHandler: &ssqHandler{Handler: basehand, service: ssqService},
	}
}
