package models

import (
	"fmt"
	"todo/config"
)

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}

func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func GetTodo(todo *Todo, id string) (err error) {
	if err = config.DB.Where("id=?", id).First(todo).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Todo, id string) (err error) {
	fmt.Println(todo)
	if err = config.DB.Where("id=?", id).Save(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *Todo, id string) (err error) {
	if err = config.DB.Where("id=?", id).Delete(todo).Error; err != nil {
		return err
	}
	return nil
}
