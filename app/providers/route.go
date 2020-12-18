package providers

import (
	"github.com/flc1125/go-gin/routes"
	"github.com/gin-gonic/gin"
)

func Route(e *gin.Engine) {
	routes.Web(e)
	routes.Api(e)
}
