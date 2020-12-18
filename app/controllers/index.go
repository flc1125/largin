package controllers

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"home":          "/",
		"user":          "/user",
		"orders.detail": "/orders/detail",
		"news.detail":   "news/detail",
	})
}
