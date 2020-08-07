package configer

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Sqlite Sqlite `json:"sqlite"`
}

type PostgreSql struct {
	Host     string // Postgresql 地址
	Port     int    // Postgresql 端口
	User     string // 数据库用户名
	Db       string // 数据库名
	Password string // 数据库用户密码
	SslMode  string // ssl 模式
}

type Sqlite struct {
	DbName string `json:"db_name"` // Postgresql 地址
	DbFile string `json:"db_file"` // Postgresql 地址
}

func NewConfiger(configPath string) *Config {
	var confTmp Config
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		log.Printf("Configuration File %v Not Exists", configPath)
	}
	readTmpFromJson(configPath, &confTmp)
	log.Printf("%v\n", confTmp)
	return &confTmp
}

//ReadTmpFromJson json 文件转map
//读取json文件转成map
func readTmpFromJson(fileName string, obj interface{}) error {
	contents, err := ioutil.ReadFile(fileName)
	if err == nil {
		er := json.Unmarshal(contents, obj)
		if er != nil {
			log.Printf("Unmarshal err:%v \n", err)
		}
	} else {
		log.Printf("read file err:%v \n", err)
	}
	return err

}
