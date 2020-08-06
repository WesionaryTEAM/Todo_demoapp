package controllers

import (
	models "Todo_demoapp/Todo_demoapp/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetTodos is ...
func GetTodos(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
