package router

import (
	"github.com/gin-gonic/gin"
	"go-web/internal/app/handler"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// 初始化gin服务
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

		v1.POST("/user", handler.CreateUser)
		v1.GET("/user", handler.ReadUser)
		v1.GET("/user/:name", handler.ReadUserByName)
	}
	return router
}
