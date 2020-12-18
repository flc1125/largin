package main

import (
	"gin.tests/app/middlewares"
	"gin.tests/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.Use(middlewares.Request())
	routes.Router(e)

	e.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"query":   "GO语言",
		})
	})

	e.Run()
}
