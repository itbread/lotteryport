package configer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Sqlite Sqlite `json:"sqlite"`
	Port   int    `json:"port"`
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
	ReadTmpFromJson(configPath, &confTmp)
	log.Printf("confTmp====%v\n", confTmp)
	return &confTmp
}

//ReadTmpFromJson json 文件转map
//读取json文件转成map
func ReadTmpFromJson(fileName string, args ...interface{}) error {
	var respStruct interface{}
	switch len(args) {
	case 1:
		respStruct = args[0]
	case 2:
		respStruct = args[1]
	default:
		return errors.New("function must have 2 or 3 arguments")
	}
	contents, err := ioutil.ReadFile(fileName)
	//log.Printf("ReadTmpFromJson err====%v contents=%v\n", err, string(contents))
	if err == nil {
		//buffer := new(bytes.Buffer)
		//buffer.Write(contents)
		//err = json.NewDecoder(buffer).Decode(&respStruct)
		err := json.Unmarshal(contents, &respStruct)
		if err != nil {
			log.Printf("Unmarshal err:%v \n", err)
		}
	} else {
		log.Printf("read file err:%v \n", err)
	}
	return err

}
