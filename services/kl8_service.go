package services

import (
	"log"

	"github.com/itbread/lotteryport/datamodels"
	"github.com/jinzhu/gorm"
)

type Kl8Service interface {
	//创建
	Create(req datamodels.Kl8) (Kl8Resp, error)
	Createlist(list []datamodels.Kl8) error
	//查询
	GetKl8(code string) (Kl8Resp, error)
	//查询list
	GetKl8s(offset int, limit int, mp map[string]interface{}) (Kl8PagingResp, error)
	//删除
	Delete(code string) error
	//更新
	Update(code string, mp map[string]interface{}) error
}

//考试评判标准 Resp
type Kl8Resp struct {
	Red1  string `json:"red1"`
	Red2  string `json:"red2"`
	Red3  string `json:"red3"`
	Red4  string `json:"red4"`
	Red5  string `json:"red5"`
	Red6  string `json:"red6"`
	Red7  string `json:"red7"`
	Red8  string `json:"red8"`
	Red9  string `json:"red9"`
	Red10 string `json:"red10"`
	Red11 string `json:"red11"`
	Red12 string `json:"red12"`
	Red13 string `json:"red13"`
	Red14 string `json:"red14"`
	Red15 string `json:"red15"`
	Red16 string `json:"red16"`
	Red17 string `json:"red17"`
	Red18 string `json:"red18"`
	Red19 string `json:"red19"`
	Red20 string `json:"red20"`
}

//分页查询
type Kl8PagingResp struct {
	Offset  int       `json:"offset"`  // Offset
	Limit   int       `json:"limit"`   // Limit
	Total   int       `json:"total"`   // Total
	Objects []SsqResp `json:"objects"` //
}

func NewKl8Service(db *gorm.DB) Kl8Service {
	svc := kl8Service{db: db}

	return &svc
}

type kl8Service struct {
	db *gorm.DB
}

func (h *kl8Service) Create(req datamodels.Kl8) (Kl8Resp, error) {
	var resp Kl8Resp
	err := h.db.Model(&datamodels.Kl8{}).
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
func (h *kl8Service) Createlist(list []datamodels.Kl8) error {
	Db := h.db.Begin()

	for _, item := range list {
		err := h.db.Model(&datamodels.Kl8{}).
			FirstOrCreate(&item, datamodels.Kl8{Code: item.Code}).Error
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

func (h *kl8Service) GetKl8(code string) (Kl8Resp, error) {
	var resp Kl8Resp
	err := h.db.Model(&datamodels.Kl8{}).
		Where("code=?", code).Scan(&resp).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *kl8Service) GetKl8s(offset int, limit int, mp map[string]interface{}) (Kl8PagingResp, error) {
	var resp Kl8PagingResp
	var err error
	resp.Offset = offset
	resp.Limit = limit

	Db := h.db.Model(&datamodels.Kl8{})
	err = Db.Where(mp).
		Order("code asc").
		Count(&resp.Total).Offset(offset).Limit(limit).
		Scan(&resp.Objects).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *kl8Service) Update(code string, mp map[string]interface{}) error {
	err := h.db.Model(datamodels.Kl8{}).
		Where("code=?", code).
		Omit("code").Updates(mp).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return err
}

func (h *kl8Service) Delete(code string) error {
	err := h.db.Where("code=?", code).
		Delete(&datamodels.Kl8{}).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return err
}
