package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Msg struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Age     int    `json:"age"`
}

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		data := map[string]interface{}{
			"name":    "czy",
			"message": "hello",
			"age":     10,
		}
		c.JSON(http.StatusOK, data)
	})

	r.GET("/json1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":    "czy",
			"message": "hello",
			"age":     10,
		})
	})
	r.Run()
}
