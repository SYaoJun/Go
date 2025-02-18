package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.GET("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "root" && password == "root" {
			c.String(http.StatusOK, "login success")
		} else {
			c.String(http.StatusBadRequest, "login failed")
		}
	})
	// http://127.0.0.1:9898/login
	r.Run(":9898")
}
