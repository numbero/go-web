package dao

import (
	"log"

	"example.com/go-web/db"
	"example.com/go-web/model"
	"gorm.io/gorm"
)

func CreateUser(user model.User) {
	db.DB.Save(&user)
	log.Println("创建用户:", user)
}

func ReadUser(name string) *model.User {
	var user model.User
	result := db.DB.Where("name = ?", name).First(&user)
	if result.Error != nil {
		log.Println("未找到用户:", name)
		return nil
	}
	log.Println("通过 name 查询用户:", user)
	return &user
}

func UpdateUser(db *gorm.DB) {
	var user model.User
	db.First(&user, 1) // 先查找 id=1 的用户

	// 更新单个字段
	db.Model(&user).Update("Name", "Bob")

	// 更新多个字段
	db.Model(&user).Updates(model.User{Name: "Charlie", Age: 35})
}

func DeleteUser(id int) {
	var user model.User
	db.DB.First(&user, id) // 找到 id=1 的用户

	// 删除用户
	db.DB.Delete(&user)
	log.Println("用户已删除")
}
