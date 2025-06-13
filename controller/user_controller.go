package controller

import (
	"log"
	"net/http"
	"strconv"

	"example.com/go-web/service"
	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	agei, _ := strconv.Atoi(age)
	service.UserAdd(name, agei)
	log.Println("添加用户:", name, agei)
	c.JSON(http.StatusOK, gin.H{
		"message": "添加成功",
	})
}
