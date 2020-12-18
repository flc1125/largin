package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		// req := c.Request

		// fmt.Println(req)

		c.Next()
	}
}
