package main

import (
	"ginweb_demo/dao"
	"ginweb_demo/models"
	"ginweb_demo/routers"
)

func main() {
	//创建数据库
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	r := routers.SetupRouter()
	r.Run(":9090")
}
