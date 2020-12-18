package payments

import "github.com/gin-gonic/gin"

func Alipay(c *gin.Context) {
	c.JSON(200, gin.H{
		"payments": "alipay",
	})
}
