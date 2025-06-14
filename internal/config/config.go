package config

import (
	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
	Options  string
}

type Config struct {
	Database DBConfig
}

var GlobalConfig Config

func InitConfig() {
	// 加载 .env 文件
	err := godotenv.Load()
	if err != nil {
		log.Println("没有找到 .env 文件，将使用系统环境变量")
	}
	appEnv := os.Getenv("APP_ENV")
	log.Println("当前环境：", appEnv)

	// 将环境变量导入配置对象
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	// 导入应用配置文件
	if _, err := toml.DecodeFile("config/config.toml", &GlobalConfig); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}
