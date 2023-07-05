package main

import (
	"github.com/gin-gonic/gin"
)

var counter = 0

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		counter++
		c.JSON(200, gin.H{
			"counter": counter,
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run()
}
