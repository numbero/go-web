package database

import (
	"fmt"
	"log"

	"go-web/internal/app/model"
	"go-web/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.GlobalConfig.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Options)
	fmt.Println("DSN:", dsn)

	var err error
	// 连接mysql数据库
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 连接sqlite数据库
	// DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	DB, err = gorm.Open(sqlite.Open("./tmp/test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	fmt.Println("✅ 数据库连接成功")

	// 自动清理数据库表
	// DB.DropTableIfExists(&User{})

	// 自动迁移模型
	err = DB.AutoMigrate(
		&model.User{},
	)
	if err != nil {
		log.Fatal("❌ 自动迁移失败:", err)
	}

	fmt.Println("✅ 数据表初始化成功")
}
