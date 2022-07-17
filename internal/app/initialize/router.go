package initialize

import (
	"github.com/gin-gonic/gin"
	"template/internal/app/middleware"
	"template/internal/app/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	//前台总路由
	frontRouter := router.FrontRouterGroupApp
	//商家总路由
	merchantRouter := router.MerchantRouterGroupApp

	PublicGroup := r.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	FrontPrivateRouter := r.Group("")
	//自定义前台中间件
	FrontPrivateRouter.Use(middleware.FrontJWTAuth())
	{
		frontRouter.User.InitFrontApiRouter(FrontPrivateRouter) //user router
	}

	MerchantPrivateRouter := r.Group("v1/merchant")
	//自定义后台中间件
	MerchantPrivateRouter.Use(middleware.MerchantJWTAuth())
	{
		merchantRouter.User.InitMerchantApiRouter(MerchantPrivateRouter) //merchant user router
	}

	return r
}
