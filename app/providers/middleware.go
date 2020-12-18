package providers

import (
	"github.com/flc1125/largin/app/middlewares"
	"github.com/gin-gonic/gin"
)

func Middleware(e *gin.Engine) {
	e.Use(middlewares.Tests())
}
