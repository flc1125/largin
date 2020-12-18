package providers

import (
	"github.com/gin-gonic/gin"
)

// 预留待定
var app = []string{
	// "database",
	"Route",
}

func Application(e *gin.Engine) {
	// 注意先后顺序
	View(e)
	Middleware(e)
	Route(e)
}
