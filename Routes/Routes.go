package routes

import (
	controllers "Todo_demoapp/Todo_demoapp/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter is ...
func SetupRouter() *gin.Engine {
	r := gin.Default()
	GroupAPI := r.Group("/GroupAPI")
	{
		GroupAPI.GET("todo", controllers.GetTodos)

	}
	return r
}
