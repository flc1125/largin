package controllers

import "github.com/gin-gonic/gin"

func OrdersDetail(c *gin.Context) {
	c.JSON(200, gin.H{
		"user": "OrdersDetail",
	})
}
