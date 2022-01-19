package routes

import (
	"github.com/gin-gonic/gin"
	"todo/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	todo := r.Group("/todo")
	{
		todo.GET("", controllers.GetTodos)
		todo.POST("", controllers.CreateTodo)

		todo.GET("/:id", controllers.GetTodo)
		todo.PUT("/:id", controllers.UpdateTodo)
		todo.DELETE("/:id", controllers.DeleteTodo)
	}
	return r
}
