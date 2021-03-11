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

type homeHandler struct {
	Handler
}

type ssqHandler struct {
	Handler
	service services.SsqService
}

type dltHandler struct {
	Handler
	service services.DltService
}
type kl8Handler struct {
	Handler
	service services.Kl8Service
}

type Handlers struct {
	SsqHandler  *ssqHandler
	DltHandler  *dltHandler
	Kl8Handler  *kl8Handler
	HomeHandler *homeHandler
}

func NewHandler(config *configer.Config) *Handlers {
	db, err := datasource.InitDb(config)
	if err != nil {
		log.Panicf("Error init data orm: %v", err)
		db.Close()
		return nil
	}
	ssqService := services.NewSsqService(db)
	dltService := services.NewDltService(db)
	kl8Service := services.NewKl8Service(db)
	basehand := Handler{config: config, db: db}
	return &Handlers{
		SsqHandler:  &ssqHandler{Handler: basehand, service: ssqService},
		DltHandler:  &dltHandler{Handler: basehand, service: dltService},
		Kl8Handler:  &kl8Handler{Handler: basehand, service: kl8Service},
		HomeHandler: &homeHandler{Handler: basehand},
	}
}
