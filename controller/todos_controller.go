package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mrifkiw/belajar_gin/initializer"
	"github.com/mrifkiw/belajar_gin/model"
)

func CreateTodos(c *gin.Context) {
	var todo model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Json data",
		})
	}

	initializer.DB.Create(&todo)

	c.JSON(200, todo)
}

func GetTodos(c *gin.Context) {
	var todos []model.Todo

	initializer.DB.Find(&todos)
	c.JSON(200, gin.H{
		"data": &todos,
	})
}
func GetTodoWithID(c *gin.Context) {
	var todo model.Todo

	todoId := c.Param("id")

	result := initializer.DB.First(&todo, todoId)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"error":      "Todo not found",
			"stackTrace": result.Error,
		})
		return
	}

	c.JSON(200, todo)
}
func UpdateTodoWithID(c *gin.Context) {
	var todo model.Todo

	todoId := c.Param("id")

	result := initializer.DB.First(&todo, todoId)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	var updateTodo model.Todo
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.JSON(400, gin.H{"error": "Invalid json data"})
		return
	}

	todo.Title = updateTodo.Title
	todo.Description = updateTodo.Description
	initializer.DB.Save(&todo)

	c.JSON(200, &todo)
}

func DeleteTodoWithID(c *gin.Context) {
	var todo model.Todo

	todoId := c.Param("id")

	result := initializer.DB.First(&todo, todoId)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	initializer.DB.Delete(&todo)

	c.JSON(200, gin.H{"message": fmt.Sprintf("Todo with id %s Deleted", todoId)})
}
