package toml

import (
	"github.com/BurntSushi/toml"

	"log"
)

type GlobalConfig struct {
	DataBase *mysql `toml:"db"`
}

var filePath = "./docs/toml/config.toml"

var globalConfig GlobalConfig

type mysql struct {
	Host        string `toml:"host"`
	User        string `toml:"user"`
	Pwd         string `toml:"password"`
	Database    string `toml:"database"`
	Type        string `toml:"type"`
	TablePrefix string `toml:"tablePrefix"`
}

func init() {
	if _, err := toml.DecodeFile(filePath, &globalConfig); err != nil {
		log.Fatal(err)
	}
}

//获取配置文件
func GetMysqlConfig() *mysql {
	return globalConfig.DataBase
}
