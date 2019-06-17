package datasource

import (
	"github.com/itbread/lotteryport/datamodels"
	"github.com/jinzhu/gorm"
)

func init()  {
	db, err := gorm.Open("sqlite3", "lotteryport.db")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&datamodels.Ssq{})
}