package app

import (
	"github.com/gin-gonic/gin"

	"gin-cli/internal/app/router"
)

func InitEngine(r router.IRouter) *gin.Engine {

	app := gin.New()

	// 配置公共中间件

	r.Register(app)


	return app
}
