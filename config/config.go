package config

import (
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

type MySQLConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	DBName    string `mapstructure:"dbname"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type AppConfig struct {
	Env    string       `mapstructure:"env"`
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	Debug  bool         `mapstructure:"debug"`
}

func Init(configPath string) *AppConfig {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv() // 自动读取环境变量

	// 设置默认值
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("env", "development")

	// MySQL 默认值
	viper.SetDefault("mysql.host", "localhost")
	viper.SetDefault("mysql.port", 3306)
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "")
	viper.SetDefault("mysql.dbname", "graduation_project")
	viper.SetDefault("mysql.charset", "utf8mb4")
	viper.SetDefault("mysql.parseTime", true)

	// Redis 默认值
	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", 6379)
	viper.SetDefault("redis.password", "")
	viper.SetDefault("redis.db", 0)

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to read config: " + err.Error())
	}

	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}

	return &cfg
}
