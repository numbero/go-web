package service

import (
	"example.com/go-web/dao"
	"example.com/go-web/model"
)

func CreateUser(name string, age int) {
	user := model.User{
		Name: name,
		Age:  age,
	}
	dao.CreateUser(user)
}

func ReadUser(name string) *model.User {
	return dao.ReadUser(name)
}
