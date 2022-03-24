package router

import (
	"github.com/gin-gonic/gin"
	"github.com/plancer16/qingdan/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", handler.IndexHandler)

	group := r.Group("/v1")
	{
		group.GET("/todo", handler.GetTodoListHandler)
		group.POST("/todo", handler.CreateTodoHandler)
		group.PUT("/todo/:id", handler.UpdateTodoHandler)
		group.DELETE("/todo/:id", handler.DeleteTodoHandler)
	}
	return r
}
