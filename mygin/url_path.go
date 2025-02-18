package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/query/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "query id = %s", id)
	})
	// http://localhost:9898/query/yaojun2025
	r.Run(":9898")
}
