package main

import (
    "myapi/internal/usecase"
    "myapi/internal/repository"
    "myapi/internal/infrastructure"
    "myapi/internal/config"
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)


func main() {
    cfg, err := config.LoadEnvConfigWithViper()
    if err != nil {
        log.Fatalf("Failed to load configuration from .env file using Viper: %s", err)
    }
    mysqlCfg := cfg.MySQL
    // MySQLデータベースに接続するための情報
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
    mysqlCfg.Username,
    mysqlCfg.Password,
    mysqlCfg.Host,
    mysqlCfg.Port,
    mysqlCfg.Database,
    )

	// MySQLデータベースに接続
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("データベース接続エラー:", err)
		return
	}

	// データベース接続を確認
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続エラー:", err)
		return
	}

	// NewMySQLRepositoryを呼び出してMySQLリポジトリを作成
	userRepository := repository.NewMySQLRepository(db)

    // userRepository := repository.NewUserRepository()           // ユーザーリポジトリの実装

	userUseCase := usecase.NewUserUsecase(userRepository)      // ユーザーユースケースの実装
    
	router := infrastructure.SetupRouter(userUseCase)         // Ginのルーティングをセットアップ

	router.Run(":8080")
}

