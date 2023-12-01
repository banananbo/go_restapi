package main

import (
    "myapi/internal/usecase"
    "myapi/internal/repository"
    "myapi/internal/infrastructure"
    "myapi/internal/config"
    "gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// GORMでのDB接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	// NewMySQLRepositoryを呼び出してMySQLリポジトリを作成
	userRepository := repository.NewMySQLRepository(db)

    // userRepository := repository.NewUserRepository()           // ユーザーリポジトリの実装

	userUseCase := usecase.NewUserUsecase(userRepository)      // ユーザーユースケースの実装
    
	router := infrastructure.SetupRouter(userUseCase)         // Ginのルーティングをセットアップ

	router.Run(":8080")

    // DB接続をクローズする
    sqlDB, _ := db.DB()
    defer sqlDB.Close()
}

