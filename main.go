package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrifkiw/belajar_gin/controller"
	"github.com/mrifkiw/belajar_gin/initializer"
)

func init() {
	initializer.ConnectToDB()
	initializer.LoadEnvVariables()
}

func main() {
	router := gin.Default()

	router.POST("/todos", controller.CreateTodos)
	router.GET("/todos", controller.GetTodos)
	router.GET("/todos/:id", controller.GetTodoWithID)
	router.PUT("/todos/:id", controller.UpdateTodoWithID)
	router.DELETE("/todos/:id", controller.DeleteTodoWithID)

	router.Run(os.Getenv("PORT"))

}
