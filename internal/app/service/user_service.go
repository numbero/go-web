package service

import (
	"fmt"
	"go-web/internal/app/model"
	"go-web/internal/app/repository"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func CreateUser(name string, age int) {
	user := model.User{
		Name: name,
		Age:  age,
	}
	repository.CreateUser(user)
}

func ReadUser(name string) *model.User {
	return repository.ReadUser(name)
}

func LoadUsers() {
	data, err := os.ReadFile("config/user.yaml")
	if err != nil {
		log.Fatalf("读取 YAML 失败: %v", err)
	}

	// 步骤 1：解析为 map[string]interface{}
	var raw map[string]interface{}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		log.Fatalf("初步解析失败: %v", err)
	}

	// 步骤 2：从 map 中提取 users
	usersRaw, ok := raw["users"]
	if !ok {
		log.Fatal("找不到 provinces 字段")
	}

	// 步骤 3：再次反序列化为结构体（嵌套结构）
	usersBytes, err := yaml.Marshal(usersRaw)
	if err != nil {
		log.Fatalf("Marshal provinces 失败: %v", err)
	}

	var users []model.User
	if err := yaml.Unmarshal(usersBytes, &users); err != nil {
		log.Fatalf("反序列化 provinces 失败: %v", err)
	}

	// 输出结果
	for _, user := range users {
		fmt.Printf("[USER] %s (%s)\n", user.Name, user.Age)
	}

	// 保存到数据库
	repository.CreateUsers(users)
}
