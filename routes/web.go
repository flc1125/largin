package routes

import (
	"gin.tests/app/controllers"
	"gin.tests/app/controllers/payments"
	"gin.tests/app/controllers/users"
	"github.com/gin-gonic/gin"
)

func user(e *gin.Engine) {
	e.GET("/", controllers.Index)
	e.GET("/user", users.Info)
	e.GET("/orders/detail", controllers.OrdersDetail)
	e.GET("/news/detail", controllers.NewsDetail)
	e.GET("/payments/alipay", payments.Alipay)
}

func Router(e *gin.Engine) {
	user(e)
}
