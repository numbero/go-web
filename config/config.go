package config

import (
	"fmt"

	"github.com/BurntSushi/toml"
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

func LoadConfig(path string) error {
	if _, err := toml.DecodeFile(path, &GlobalConfig); err != nil {
		return fmt.Errorf("读取配置文件失败: %w", err)
	}
	return nil
}
