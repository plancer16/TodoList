package model

import (
	"github.com/plancer16/qingdan/dao"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return err
}

func GetTodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.Find(&todo).Where("id=?", id).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodoList() (todoList []*Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return todoList, nil
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return err
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return err
}
