package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"message": "index",
	})
}

func costHandler(c *gin.Context) {
	fmt.Println("cost handler")
	start := time.Now()
	c.Next() //调用后续的处理函数
	cost := time.Since(start)
	fmt.Println("cost=", cost)
}

func main() {
	r := gin.Default()

	//全局注册中间件
	//r.Use(costHandler)

	r.GET("/index", costHandler, indexHandler)
	r.Run()
}
