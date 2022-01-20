package models

import (
	"fmt"
	"todo/config"
)

type Todo struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (b *Todo) TableName() string {
	return "todo"
}

func GetAllTodos(todo *[]Todo) (err error) {
	/* grom
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	*/
	if err = config.DB.Find(todo); err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *Todo) (err error) {
	/* gorm
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}*/
	_, err = config.DB.Insert(todo)
	if err != nil {
		return err
	}
	return nil
}

func GetTodo(todo *Todo, id string) (err error) {
	/* gorm
	if err = config.DB.Where("id=?", id).First(todo).Error; err != nil {
		return err
	}*/
	_, err = config.DB.ID(id).Get(todo)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTodo(todo *Todo, id string, params []string) (err error) {
	fmt.Println(params)
	/* gorm
	if err = config.DB.Where("id=?", id).Save(todo).Error; err != nil {
		return err
	}*/
	_, err = config.DB.ID(id).Update(todo)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *Todo, id string) (err error) {
	/* gorm
	if err = config.DB.Where("id=?", id).Delete(todo).Error; err != nil {
		return err
	}
	*/
	_, err = config.DB.ID(id).Delete(todo)
	if err != nil {
		return err
	}
	return nil
}
