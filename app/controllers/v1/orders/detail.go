package orders

import "github.com/gin-gonic/gin"

func Detail(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": 123123,
	})
}
