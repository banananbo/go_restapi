package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	router := gin.Default()

	// GETリクエストを処理するエンドポイント
	router.GET("/users", ListUsers)

	// POSTリクエストを処理するエンドポイント
	router.POST("/users/add", AddUser)

	router.Run(":8080")
}

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func AddUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser.ID = "some_generated_id" // 本当はここでIDを生成するロジックを実装する

	users = append(users, newUser)
	c.JSON(http.StatusOK, gin.H{"status": "User added successfully", "user": newUser})
}