package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Port         string `ini:"port"`
	Release      string `ini:"release"`
	*MysqlConfig `ini:"mysql"`
}

type MysqlConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Db       string `ini:"db"`
}

func Init(file string) (err error) {
	return ini.MapTo(Conf, file)
}
