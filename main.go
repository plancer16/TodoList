package main

import (
	"fmt"
	"github.com/plancer16/qingdan/dao"
	"github.com/plancer16/qingdan/model"
	"github.com/plancer16/qingdan/router"
	"github.com/plancer16/qingdan/setting"
)

func main() {
	err := setting.Init("conf/config.ini")
	if err != nil {
		fmt.Sprintf("load config failed, err: %s", err.Error())
		return
	}
	err = dao.InitDB(setting.Conf.MysqlConfig)
	if err != nil {
		fmt.Sprintf("init db failed, err:%s", err.Error())
		return
	}
	dao.DB.AutoMigrate(&model.Todo{})
	defer dao.DB.Close()

	r := router.SetupRouter()
	r.Run()
}
