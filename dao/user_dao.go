package dao

import (
	"log"

	"example.com/go-web/db"
	"example.com/go-web/model"
	"gorm.io/gorm"
)

func AddUser(user model.User) {
	db.DB.Save(&user)
}

func CreateUser(user model.User) {
	db.DB.Save(&user)
	log.Println("创建用户:", user)
}

func ReadUser(id int) {
	var user model.User

	// 查询主键为 1 的用户
	db.DB.First(&user, id)
	log.Println("查询到用户:", user)

	// 查询某个 email 的用户
	db.DB.Where("email = ?", "alice@example.com").First(&user)
	log.Println("通过 email 查询用户:", user)
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
