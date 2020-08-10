package datamodels

type Ssq struct {
	//gorm.Model
	Code string `gorm:"column:code;primary_key;type:varchar(25)"` // 文件唯一标识
	Date string `json:"date"`
	Week string `json:"week"`
	BallInfo
}

type BallInfo struct {
	Red1 string `json:"red1"`
	Red2 string `json:"red2"`
	Red3 string `json:"red3"`
	Red4 string `json:"red4"`
	Red5 string `json:"red5"`
	Red6 string `json:"red6"`
	Red  string `json:"red"`
	Blue string `json:"blue"`
}
