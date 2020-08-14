package datasource

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/itbread/lotteryport/configer"
	"github.com/itbread/lotteryport/datamodels"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDb(config *configer.Config) (db *gorm.DB, err error) {
	fmt.Printf("config:%v", spew.Sprintf("%#v", config))
	db, err = gorm.Open(config.Sqlite.DbName, config.Sqlite.DbFile)
	if err != nil {
		panic("连接数据库失败")
		return nil, err
	}
	//defer db.Close()
	db.DB().SetConnMaxLifetime(240)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10) //解决 有时返回 EOF
	db.SingularTable(true)

	// 自动迁移模式
	db.AutoMigrate(&datamodels.Ssq{})
	db.AutoMigrate(&datamodels.Dlt{})
	return db, err
}
