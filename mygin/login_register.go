package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	user := r.Group("/user")
	user.GET("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "root" && password == "root" {
			c.String(http.StatusOK, "login success")
		} else {
			c.String(http.StatusBadRequest, "login failed")
		}
	})
	user.GET("/register", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		// 1. 检测密码格式
		// 2. 存储到数据库
		password_length := len(password)
		username_length := len(username)
		if password_length < 8 || username_length > 8 {
			c.String(http.StatusBadRequest, "password too short or username too long")
		} else {
			c.String(http.StatusOK, "register success")
		}
	})
	// http://127.0.0.1:9898/user/login
	// http://127.0.0.1:9898/user/register
	r.Run(":9898")
}
