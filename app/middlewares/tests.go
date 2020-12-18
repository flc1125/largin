package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Tests() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.Request

		fmt.Println(req)

		c.Next()
	}
}
