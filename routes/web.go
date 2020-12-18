package routes

import (
	"github.com/flc1125/go-gin/app/controllers"
	"github.com/flc1125/go-gin/app/controllers/payments"
	"github.com/flc1125/go-gin/app/controllers/users"
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
