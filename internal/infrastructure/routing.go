package infrastructure

import (
	"github.com/gin-gonic/gin"
	"myapi/internal/usecase"
	"myapi/internal/domain"
	"net/http"
)


// SetupRouter はGinのルーティングを設定します
func SetupRouter(userUC usecase.UserUsecase) *gin.Engine {
	router := gin.Default()

	// ユーザー関連のエンドポイント
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", func(c *gin.Context) {
			userList, err := userUC.GetUserList()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"users": userList})
		})

		userGroup.POST("/add", func(c *gin.Context) {
			// 新しいユーザーを追加するハンドラ
			// userUC.AddUser() を呼び出してデータを追加する
			var newUser domain.User
			if err := c.ShouldBindJSON(&newUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			addedUser, err := userUC.AddUser(newUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"status": "User added successfully", "user": addedUser})
		})
	}

	// 他のエンドポイントを追加...

	return router
}
