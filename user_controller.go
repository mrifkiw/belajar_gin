package main

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc *UserController) getUserInfo(c *gin.Context) {
	userID := c.Param("id")
	// Fetch user information from the database or other data source
	// For simplicity, we'll just return a JSON response.
	c.JSON(200, gin.H{
		"id":    userID,
		"name":  "John Doe",
		"email": "john@example.com",
	})
}
