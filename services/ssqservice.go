package services

import (
	"github.com/itbread/lotteryport/datamodels"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SsqService interface {
	GetAll() []datamodels.Ssq
	GetByID(id int64) (datamodels.Ssq, bool)
	DeleteByID(id int64) bool
	Update(id int64, user datamodels.Ssq) (datamodels.Ssq, error)
	Create(red string, ssq datamodels.Ssq) (datamodels.Ssq, error)
}


