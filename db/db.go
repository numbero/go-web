package db

import (
	"fmt"
	"log"

	"example.com/go-web/config"
	"example.com/go-web/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.GlobalConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Options)
	fmt.Println("DSN:", dsn)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	fmt.Println("✅ 数据库连接成功")

	// 自动迁移模型
	err = DB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal("❌ 自动迁移失败:", err)
	}

	fmt.Println("✅ 数据表初始化成功")
}
