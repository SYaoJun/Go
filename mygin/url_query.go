package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/query", func(c *gin.Context) {
		q := c.Query("q")
		c.String(http.StatusOK, "query value = %s", q)
	})
	// http://127.0.0.1:9898/query?q=book
	r.Run(":9898")
}
