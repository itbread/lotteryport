package datamodels

type Kl8 struct {
	Code string `gorm:"column:code;primary_key;type:varchar(25)"` // 文件唯一标识
	Date string `json:"date"`
	Week string `json:"week"`
	Red  string `json:"red"`
	KlBallInfo
}

type KlBallInfo struct {
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

type KlBallInfo2 struct {
	Red1  string `gorm:"column:red1"`
	Red2  string `gorm:"column:red2"`
	Red3  string `gorm:"column:red3"`
	Red4  string `gorm:"column:red4"`
	Red5  string `gorm:"column:red5"`
	Red6  string `gorm:"column:red6"`
	Red7  string `gorm:"column:red7"`
	Red8  string `gorm:"column:red8"`
	Red9  string `gorm:"column:red9"`
	Red10 string `gorm:"column:red10"`
	Red11 string `gorm:"column:red11"`
	Red12 string `gorm:"column:red12"`
	Red13 string `gorm:"column:red13"`
	Red14 string `gorm:"column:red14"`
	Red15 string `gorm:"column:red15"`
	Red16 string `gorm:"column:red16"`
	Red17 string `gorm:"column:red17"`
	Red18 string `gorm:"column:red18"`
	Red19 string `gorm:"column:red19"`
	Red20 string `gorm:"column:red20"`
}
