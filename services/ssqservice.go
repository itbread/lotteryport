package services

import (
	"github.com/itbread/lotteryport/datamodels"
	"github.com/jinzhu/gorm"
	"log"
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
	SerialId uint64 `json:"serial_id"` // 扣分项序号
	Type     int    `json:"type"`      // 所属项目 0通用要求 1倒车 2坡道 3侧方 4曲线 5直角 6项目间道路
	ItemDesc string `json:"item_desc"` // 扣分项目说明
	Grade    int    `json:"grade"`     // 扣分分数
	CreateDt int64  `json:"create_dt"` // 创建日期
	UpdateDt int64  `json:"update_dt"` // 更新日期
	KmType   int    `json:"km_type"`   // 科目类型 1 :k2 2 :k3
	PlaceId  string `json:"place_id"`  // 场地编号
}

//分页查询
type SsqPagingResp struct {
	Offset  int       `json:"offset"`  // Offset
	Limit   int       `json:"limit"`   // Limit
	Total   int       `json:"total"`   // Total
	Objects []SsqResp `json:"objects"` //
}

func NewSsqService(db *gorm.DB) SsqService {
	svc := examStandardService{db: db}

	return &svc
}

type examStandardService struct {
	db *gorm.DB
}

func (h *examStandardService) Create(req datamodels.Ssq) (SsqResp, error) {
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
func (h *examStandardService) Createlist(list []datamodels.Ssq) error {
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

func (h *examStandardService) GetSsq(code string) (SsqResp, error) {
	var resp SsqResp
	err := h.db.Model(&datamodels.Ssq{}).
		Where("code=?", code).Scan(&resp).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *examStandardService) GetSsqs(offset int, limit int, mp map[string]interface{}) (SsqPagingResp, error) {
	var resp SsqPagingResp
	var err error
	resp.Offset = offset
	resp.Limit = limit

	Db := h.db.Model(&datamodels.Ssq{})
	err = Db.Where(mp).
		Order("serial_id asc").
		Count(&resp.Total).Offset(offset).Limit(limit).
		Scan(&resp.Objects).Error

	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return resp, err
}

func (h *examStandardService) Update(code string, mp map[string]interface{}) error {
	err := h.db.Model(datamodels.Ssq{}).
		Where("code=?", code).
		Omit("code").Updates(mp).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	return err
}

func (h *examStandardService) Delete(code string) error {
	err := h.db.Where("code=?", code).
		Delete(&datamodels.Ssq{}).Error
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	return err
}
