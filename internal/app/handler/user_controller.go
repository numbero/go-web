package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-web/internal/app/service"
)

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	agei, _ := strconv.Atoi(age)
	service.CreateUser(name, agei)
	log.Println("添加用户:", name, agei)
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}

func ReadUser(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name is required",
		})
		return
	}
	user := service.ReadUser(name)
	log.Println("查询用户:", name)
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"data":    user,
	})
}

func ReadUserByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "name is required",
		})
		return
	}
	user := service.ReadUser(name)
	log.Println("查询用户:", name)
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"data":    user,
	})
}
