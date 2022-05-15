package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{"method": "get"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "post"})
		}
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	})

	group := r.Group("/video")
	{
		group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "index",
			})
		})
		group.GET("/aa", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"method": "aa",
			})
		})
	}

	r.Run()
}
