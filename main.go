package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_gin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Todo{})

	todosController := &TodosController{}

	// db.Create(&Todo{title: "hello", description: "description"})

	router.POST("/todos", func(ctx *gin.Context) {
		var todo Todo
		if err := ctx.ShouldBindJSON(&todo); err != nil {
			ctx.JSON(400, gin.H{
				"error": "Invalid Json data",
			})
		}

		db.Create(&todo)

		ctx.JSON(200, todo)
	})

	router.GET("/todos", func(ctx *gin.Context) {
		var todos []Todo

		db.Find(&todos)
		ctx.JSON(200, todos)
	})

	router.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo

		todoId := c.Param("id")

		result := db.First(&todo, todoId)

		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		c.JSON(200, todo)
	})
	router.PUT("/todos/:id", func(c *gin.Context) {
		var todo Todo

		todoId := c.Param("id")

		result := db.First(&todo, todoId)

		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		var updateTodo Todo
		if err := c.ShouldBindJSON(&updateTodo); err != nil {
			c.JSON(400, gin.H{"error": "Invalid json data"})
			return
		}

		todo.Title = updateTodo.Title
		todo.Description = updateTodo.Description
		db.Save(&todo)

		c.JSON(200, &todo)
	})
	router.DELETE("/todos/:id", func(c *gin.Context) {
		var todo Todo

		todoId := c.Param("id")

		result := db.First(&todo, todoId)

		if result.Error != nil {
			c.JSON(404, gin.H{"error": "Todo not found"})
			return
		}

		db.Delete(&todo)

		c.JSON(200, gin.H{"message": fmt.Sprintf("Todo with id %s Deleted", todoId)})
	})

	router.Run(":8000")

}
