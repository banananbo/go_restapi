// config/loader.go

package config

import (
    "github.com/spf13/viper"
)

// LoadEnvConfigWithViper はViperを使って.envファイルを読み込んで設定情報を返します
func LoadEnvConfigWithViper() (AppConfig, error) {
    var cfg AppConfig

    viper.SetConfigFile(".env")
    err := viper.ReadInConfig()
    if err != nil {
        return cfg, err
    }

    cfg.MySQL.Username = viper.GetString("MYSQL_USERNAME")
    cfg.MySQL.Password = viper.GetString("MYSQL_PASSWORD")
    cfg.MySQL.Host = viper.GetString("MYSQL_HOST")
    cfg.MySQL.Port = viper.GetString("MYSQL_PORT")
    cfg.MySQL.Database = viper.GetString("MYSQL_DATABASE")

    return cfg, nil
}
