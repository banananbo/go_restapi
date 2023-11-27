package main

import (
    "myapi/internal/usecase"
    "myapi/internal/repository"
    "myapi/internal/infrastructure"
)


func main() {
    userRepository := repository.NewUserRepository()           // ユーザーリポジトリの実装
	userUseCase := usecase.NewUserUsecase(userRepository)      // ユーザーユースケースの実装
    
	router := infrastructure.SetupRouter(userUseCase)         // Ginのルーティングをセットアップ

	router.Run(":8080")
}

