package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TodosController struct{}

func (tc *TodosController) createTodos(c *gin.Context, db gorm.DB) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Json data",
		})
	}

	db.Create(&todo)

	c.JSON(200, todo)
}
