package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/plancer16/qingdan/setting"
)

var DB *gorm.DB

func InitDB(config *setting.MysqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", config.User, config.Password, config.Host, config.Port, config.Db)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}
