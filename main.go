package main

import (
	"github.com/flc1125/go-gin/app/providers"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	providers.Application(e)

	e.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
			"query":   "GO语言",
		})
	})

	e.Run()
}
