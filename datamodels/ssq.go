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


type SsqBallInfo struct {
	State     int    `json:"state"`
	Message   string `json:"message"`
	PageCount int    `json:"pageCount"`
	CountNum  int    `json:"countNum"`
	Tflag     int    `json:"Tflag"`
	Result    []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		DetailsLink string `json:"detailsLink"`
		VideoLink   string `json:"videoLink"`
		Date        string `json:"date"`
		Week        string `json:"week"`
		Red         string `json:"red"`
		Blue        string `json:"blue"`
		Blue2       string `json:"blue2"`
		Sales       string `json:"sales"`
		Poolmoney   string `json:"poolmoney"`
		Content     string `json:"content"`
		Addmoney    string `json:"addmoney"`
		Addmoney2   string `json:"addmoney2"`
		Msg         string `json:"msg"`
		Z2Add       string `json:"z2add"`
		M2Add       string `json:"m2add"`
		Prizegrades []struct {
			Type      int    `json:"type"`
			Typenum   string `json:"typenum"`
			Typemoney string `json:"typemoney"`
		} `json:"prizegrades"`
	} `json:"result"`
}