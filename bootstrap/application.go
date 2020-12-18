package bootstrap

import (
	"github.com/flc1125/go-gin/app/providers"
	"github.com/gin-gonic/gin"
)

func Application() *gin.Engine {
	e := gin.Default()

	// 服务，注意先后顺序（如：模板、中间件、路由）
	providers.App(e)        // 全局应用服务
	providers.Cache(e)      // 缓存
	providers.View(e)       // 启用模板
	providers.Middleware(e) // 中间件
	providers.Route(e)      // 路由

	return e
}
