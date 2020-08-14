package services

import (
	"log"

	"github.com/itbread/lotteryport/datamodels"
	"github.com/jinzhu/gorm"
)

type SsqService interface {
	//创建
	Create(req datamodels.Ssq) (SsqResp, error)
	Createlist(list []datamodels.Ssq) error
	//查询
	GetSsq(code string) (SsqResp, error)
	//查询list
	GetSsqs(offset int, limit int, mp map[string]interface{}) (SsqPagingResp, error)
	//删除
	Delete(code string) error
	//更新
	Update(code string, mp map[string]interface{}) error
}

//考试评判标准 Resp
type SsqResp struct {
	Red1  string `json:"red1"`
	Red2  string `json:"red2"`
	Red3  string `json:"red3"`
	Red4  string `json:"red4"`
	Red5  string `json:"red5"`
	Red6  string `json:"red6"`
	Blue1 string `json:"red"`
	Blue2 string `json:"blue"`
}

//分页查询
type SsqPagingResp struct {
	Offset  int       `json:"offset"`  // Offset
	Limit   int       `json:"limit"`   // Limit
	Total   int       `json:"total"`   // Total
	Objects []SsqResp `json:"objects"` //
}

func NewSsqService(db *gorm.DB) SsqService {
	svc := ssqService{db: db}

	return &svc
}

type ssqService struct {
	db *gorm.DB
}

func (h *ssqService) Create(req datamodels.Ssq) (SsqResp, error) {
	var resp SsqResp
	err := h.db.Model(&datamodels.Ssq{}).
		Create(&req).Error

	if err != nil {
		log.Printf("Create error: %v\n", err)
		return resp, err
	}

	if err != nil {
		log.Printf("StructCopy error: %v\n", err)
	}
	return resp, err

}
func (h *ssqService) Createlist(list []datamodels.Ssq) error {
	Db := h.db.Begin()

	for _, item := range list {
		err := h.db.Model(&datamodels.Ssq{}).
			FirstOrCreate(&item, datamodels.Ssq{Code: item.Code}).Error
		if err != nil {
			Db.Rollback()
			log.Printf("Createlist error: %v\n", err)
			return err
		}
	}
	err := Db.Commit().Error

	if err != nil {
		Db.Rollback()
		log.Printf("Db.Commit().Error: %v\n", err)
		return err
	}

	return err

}

func (h *ssqService) GetSsq(code string) (SsqResp, error) {
	var resp SsqResp
	err := h.db.Model(&datamodels.Ssq{}).
		Where("code=?", code).Scan(&resp).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *ssqService) GetSsqs(offset int, limit int, mp map[string]interface{}) (SsqPagingResp, error) {
	var resp SsqPagingResp
	var err error
	resp.Offset = offset
	resp.Limit = limit

	Db := h.db.Model(&datamodels.Ssq{})
	err = Db.Where(mp).
		Order("code asc").
		Count(&resp.Total).Offset(offset).Limit(limit).
		Scan(&resp.Objects).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *ssqService) Update(code string, mp map[string]interface{}) error {
	err := h.db.Model(datamodels.Ssq{}).
		Where("code=?", code).
		Omit("code").Updates(mp).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return err
}

func (h *ssqService) Delete(code string) error {
	err := h.db.Where("code=?", code).
		Delete(&datamodels.Ssq{}).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return err
}
