package main

import "github.com/gin-gonic/gin"

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required,min=18"`
}

func createUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
func main() {
	r := gin.Default()
	r.POST("/users", createUser)
	r.Run(":8080")
}
