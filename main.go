package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"go-web/internal/app/model"
	"go-web/internal/app/router"
	"go-web/internal/config"
	"go-web/internal/database"
	"go-web/internal/logger"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func main() {
	// 加载配置文件 (导入逻辑已经转移到config包中)
	config.InitConfig()
	log.Println("加载配置文件成功")

	// 初始化日志
	logger.Init()
	log.Println("初始化日志成功")

	// 初始化数据库连接
	database.InitDB()
	log.Println("初始化数据库表成功")

	// 测试
	//service.LoadAddresses()
	//service.LoadUsers()
	demo_json()

	// 初始化路由
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("启动服务器失败:", err)
	}
	log.Println("服务器启动成功")
}

func demo_json() {
	jsonStr := `{
        "user": {
            "profile": {
                "name": "Alice",
                "age": 30
            },
            "hobbies": ["reading", "cycling"]
        }
    }`

	// 链式获取深层字段
	name := jsoniter.Get([]byte(jsonStr), "user", "profile", "name").ToString()
	age := jsoniter.Get([]byte(jsonStr), "user", "profile", "age").ToInt()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	// 获取数组字段
	hobbies := jsoniter.Get([]byte(jsonStr), "user", "hobbies")
	fmt.Print("Hobbies: ")
	for i := 0; i < hobbies.Size(); i++ {
		fmt.Print(hobbies.Get(i).ToString(), " ")
	}
	fmt.Println()

	// 自定义测试反序列化
	data, err := os.ReadFile("config/user.yaml")
	if err != nil {
		log.Fatalf("读取 YAML 失败: %v", err)
	}
	// yaml反序列化为 map[string]interface{}
	var userMap map[string]interface{}
	if err := yaml.Unmarshal(data, &userMap); err != nil {
		log.Fatalf("初步解析失败: %v", err)
	}
	// json序列化
	jsonBytes, err := jsoniter.Marshal(userMap)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	rootNode := jsoniter.Get(jsonBytes)
	fmt.Println(rootNode.Get("version").ToString())
	var users []model.User
	rootNode.Get("users").ToVal(&users)
	for _, user := range users {
		fmt.Println(user.Name, user.Age)
	}

	usersJson, err := jsoniter.Marshal(users)
	if err != nil {
		fmt.Println("序列化失败", err)
	}
	fmt.Println("序列化结果：", string(usersJson))
}
