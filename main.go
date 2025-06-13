package main

import (
	"log"
	"net/http"

	"example.com/go-web/config"
	"example.com/go-web/controller"
	"example.com/go-web/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载 TOML 配置文件
	if err := config.LoadConfig("config/config.toml"); err != nil {
		log.Fatal(err)
	}
	log.Println("加载配置文件成功")

	// 初始化数据库连接
	db.InitDB()
	log.Println("初始化数据库表成功")

	// 初始化gin服务
	// e := gin.Default()

	// e.GET("/user/add", controller.UserAdd)

	// if err := e.Run(":18080"); err != nil {
	// 	log.Fatal("启动服务器失败:", err)
	// }

	// 1. 创建一个默认的 Gin 引擎
	// gin.Default() 会返回一个包含了 Logger 和 Recovery 中间件的 Gin 引擎
	router := gin.Default()

	// 2. 定义一个 GET 路由
	// 当客户端以 GET 方法请求 /ping 路径时，会执行后面的匿名函数
	router.GET("/ping", func(c *gin.Context) {
		// c.JSON() 是一个辅助函数，可以方便地返回 JSON 响应
		// 它接收状态码和任意可以被序列化为 JSON 的对象
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// 定义路由组
	v1 := router.Group("/api/v1")
	{
		// v1.POST("/", createTodo)
		// v1.GET("/", fetchAllTodo)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)

		v1.POST("/user/add", controller.UserAdd)
	}

	// 3. 启动 HTTP 服务器
	// 默认监听在 :8080 端口
	// router.Run() 在内部调用了 http.ListenAndServe
	err := router.Run(":8080")
	if err != nil {
		// 如果服务器启动失败，这里可以处理错误
		// 例如：log.Fatal("Failed to start server: ", err)
		return
	}
}
