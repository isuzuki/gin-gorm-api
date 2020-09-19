package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/item", func(c *gin.Context) {
		item := &Item{1, "item1", "desc", "hidden1", "hidden2"}
		c.JSON(http.StatusOK, item)
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
