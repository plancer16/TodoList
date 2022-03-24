package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/plancer16/qingdan/model"
	"net/http"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func GetTodoListHandler(ctx *gin.Context) {
	list, err := model.GetTodoList()
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, list)
}

func CreateTodoHandler(ctx *gin.Context) {
	var todo model.Todo
	err := ctx.BindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	err = model.CreateTodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func UpdateTodoHandler(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"err": "无效id"})
		return
	}
	todo, err := model.GetTodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	err = ctx.BindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	err = model.UpdateTodo(todo)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, todo)
}

func DeleteTodoHandler(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{"err": "无效id"})
		return
	}
	err := model.DeleteTodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{id: "deleted"})
}
