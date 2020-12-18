package providers

import "github.com/gin-gonic/gin"

func View(e *gin.Engine) {
	e.LoadHTMLGlob("resources/views/*")
}
