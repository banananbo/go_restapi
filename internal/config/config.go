// config/config.go

package config

// MySQLConfig はMySQLの接続情報を保持する構造体です
type MySQLConfig struct {
    Username string
    Password string
    Host     string
    Port     string
    Database string
    // 他のMySQL接続に関する設定項目を追加できます
}

// AppConfig はアプリケーション全体の設定情報を保持する構造体です
type AppConfig struct {
    MySQL MySQLConfig
    // 他のアプリケーション設定項目を追加できます
}
