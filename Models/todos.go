package models

import (
	config "Todo_demoapp/Todo_demoapp/Config"
)

//GetAllTodos is ...
func GetAllTodos(todo *[]Todo) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}
