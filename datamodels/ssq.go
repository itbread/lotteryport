package datamodels

type Ssq struct {
	//gorm.Model
	Code string `gorm:"unique" json:"code"`
	Date string  `json:"date"`
	Week string  `json:"week"`
	Red  string  `json:"red"`
	Blue string  `json:"blue"`
}
