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
	r.GET("/items", func(c *gin.Context) {
		type ItemCount struct {
			Count int `json:"count"`
			Item
		}
		item1 := Item{1, "item1", "desc", "hidden1", "hidden2"}
		item2 := Item{2, "item2", "desc", "hidden1", "hidden2"}

		var itemCounts []ItemCount
		itemCounts = append(itemCounts, ItemCount{10, item1})
		itemCounts = append(itemCounts, ItemCount{12, item2})
		c.JSON(http.StatusOK, gin.H{
			"items": itemCounts,
		})
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
