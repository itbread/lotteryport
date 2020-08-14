package datamodels

type Dlt struct {
	//gorm.Model
	Code string `gorm:"column:code;primary_key;type:varchar(25)";json:"code"` // 文件唯一标识
	Date string `json:"date"`
	Week string `json:"week"`
	DltBallInfo
}

type DltBallInfo struct {
	Red1  string `json:"red1"`
	Red2  string `json:"red2"`
	Red3  string `json:"red3"`
	Red4  string `json:"red4"`
	Red5  string `json:"red5"`
	Red6  string `json:"red6"`
	Blue1 string `json:"red"`
	Blue2 string `json:"blue"`
}
