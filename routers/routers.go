package routers

import (
	"ginweb_demo/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)

		//查看
		//查看所有
		v1Group.GET("/todo", controller.GetTodoList)
		//查看某一个
		//...

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}