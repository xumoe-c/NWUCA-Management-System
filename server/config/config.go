package config

import (
	"github.com/spf13/viper"
)

// Config 结构体用于映射所有配置
type Config struct {
	Server ServerConfig
	DB     DatabaseConfig
	JWT    JWTConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// JWTConfig JWT 配置
type JWTConfig struct {
	SecretKey      string
	ExpirationDays int
}

// LoadConfig 从文件或环境变量加载配置
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")   // 配置文件名 (无扩展名)
	viper.SetConfigType("yaml")     // 配置文件类型
	viper.AddConfigPath("./config") // 配置文件路径
	viper.AutomaticEnv()            // 允许从环境变量读取

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
