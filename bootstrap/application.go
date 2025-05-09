package bootstrap

import (
	"github.com/flc1125/largin/app/providers"
	"github.com/gin-gonic/gin"
)

func Application() *gin.Engine {
	e := gin.Default()

	// 服务管理，注意先后顺序
	providers.Env(e)        // 环境配置
	providers.Config(e)     // 系统配置
	providers.App(e)        // 全局应用服务
	providers.Cache(e)      // 缓存
	providers.View(e)       // 模板（若无需模板，可注释）
	providers.Middleware(e) // 中间件
	providers.Route(e)      // 路由

	return e
}
