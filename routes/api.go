package routes

import "github.com/gin-gonic/gin"

func Api(e *gin.Engine) {
	e.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"api_version": "1.0.0",
		})
	})
}
