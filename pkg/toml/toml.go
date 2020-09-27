package toml

import (
	"github.com/BurntSushi/toml"
	"time"

	"log"
)

type GlobalConfig struct {
	RunMode  string  `toml:"run_mode"`
	DataBase *mysql  `toml:"db"`
	App      *app    `toml:"app"`
	Server   *server `toml:"server"`
}

var filePath = "./docs/toml/config.toml"

var globalConfig GlobalConfig

type app struct {
	PageSize  int    `toml:"page_size"`
	JwtSecret string `toml:"jwt_secret"`
}

type server struct {
	HttpPort        int `toml:"http_port"`
	ReadTimeoutInt  int `toml:"read_timeout"`
	WriteTimeoutInt int `toml:"write_timeout"`
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

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

//获取配置文件
func GetAppConfig() *app {
	return globalConfig.App
}

//获取配置文件
func GetServerConfig() *server {
	server := globalConfig.Server
	server.ReadTimeout = time.Duration(server.ReadTimeoutInt) * time.Second
	server.WriteTimeout = time.Duration(server.WriteTimeoutInt) * time.Second
	return globalConfig.Server
}

func GetRunMode() string {
	return globalConfig.RunMode
}
