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
		GroupAPI.POST("todo", controllers.CreateATodo)
		GroupAPI.GET("todo/:id", controllers.GetATodo)
		GroupAPI.PUT("todo/:id", controllers.UpdateATodo)
		GroupAPI.DELETE("todo/:id", controllers.DeleteATodo)
	}
	return r
}
