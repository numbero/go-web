package service

import (
	"log"

	"example.com/go-web/dao"
	"example.com/go-web/model"
)

func UserAdd(name string, age int) {
	user := model.User{
		Name: name,
		Age:  age,
	}
	log.Println("准备添加用户:", user)

	dao.AddUser(user)
}
