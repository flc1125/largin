package routes

import (
	"github.com/flc1125/largin/app/controllers"
	"github.com/flc1125/largin/app/controllers/payments"
	"github.com/flc1125/largin/app/controllers/users"
	v1_orders "github.com/flc1125/largin/app/controllers/v1/orders"
	"github.com/gin-gonic/gin"
)

func user(e *gin.Engine) {
	e.GET("/", controllers.Welcome)
	e.GET("/user", users.Info)
	e.GET("/payments/alipay", payments.Alipay)
}

func group(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		v1.GET("/orders/detail", v1_orders.Detail)

		v1.GET("/version", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"version": "v1",
			})
		})
	}
}

func Web(e *gin.Engine) {
	user(e)
	group(e)
}
