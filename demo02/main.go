package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		name := c.Query("name")
		query := c.DefaultQuery("query", "default")
		c.JSON(http.StatusOK, gin.H{
			"result": "success",
			"name":   name,
			"query":  query,
		})
	})
	r.Run()
}
