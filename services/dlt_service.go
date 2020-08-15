package services

import (
	"log"

	"github.com/itbread/lotteryport/datamodels"
	"github.com/jinzhu/gorm"
)

type DltService interface {
	//创建
	Create(req datamodels.Dlt) (DltResp, error)
	Createlist(list []datamodels.Dlt) error
	//查询
	GetDlt(code string) (DltResp, error)
	//查询list
	GetDlts(offset int, limit int, mp map[string]interface{}) (DltPagingResp, error)
	//删除
	Delete(code string) error
	//更新
	Update(code string, mp map[string]interface{}) error
}

//考试评判标准 Resp
type DltResp struct {
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
type DltPagingResp struct {
	Offset  int       `json:"offset"`  // Offset
	Limit   int       `json:"limit"`   // Limit
	Total   int       `json:"total"`   // Total
	Objects []SsqResp `json:"objects"` //
}

func NewDltService(db *gorm.DB) DltService {
	svc := dltService{db: db}

	return &svc
}

type dltService struct {
	db *gorm.DB
}

func (h *dltService) Create(req datamodels.Dlt) (DltResp, error) {
	var resp DltResp
	err := h.db.Model(&datamodels.Dlt{}).
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
func (h *dltService) Createlist(list []datamodels.Dlt) error {
	Db := h.db.Begin()

	for _, item := range list {
		err := h.db.Model(&datamodels.Dlt{}).
			FirstOrCreate(&item, datamodels.Dlt{Code: item.Code}).Error
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

func (h *dltService) GetDlt(code string) (DltResp, error) {
	var resp DltResp
	err := h.db.Model(&datamodels.Dlt{}).
		Where("code=?", code).Scan(&resp).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *dltService) GetDlts(offset int, limit int, mp map[string]interface{}) (DltPagingResp, error) {
	var resp DltPagingResp
	var err error
	resp.Offset = offset
	resp.Limit = limit

	Db := h.db.Model(&datamodels.Dlt{})
	err = Db.Where(mp).
		Order("code asc").
		Count(&resp.Total).Offset(offset).Limit(limit).
		Scan(&resp.Objects).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *dltService) Update(code string, mp map[string]interface{}) error {
	err := h.db.Model(datamodels.Dlt{}).
		Where("code=?", code).
		Omit("code").Updates(mp).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return err
}

func (h *dltService) Delete(code string) error {
	err := h.db.Where("code=?", code).
		Delete(&datamodels.Dlt{}).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return err
}
