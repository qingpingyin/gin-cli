package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"template/internal/app/api/v1/user"
	"template/internal/app/middleware"
)

var RouterProviderSet = wire.NewSet(wire.Struct(new(Router), "*"))

type IRouter interface {
	InitRouter(router *gin.RouterGroup)
}

type Router struct {
	UserAPI *user.UserApi
}

func (a *Router) RegisterAPI() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")

	publicGroup := v1.Group("")
	{
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	frontPrivateRouter := r.Group("front")
	frontPrivateRouter.Use(middleware.FrontJWTAuth())
	{
		a.UserAPI.FrontUserApi.InitRouter(frontPrivateRouter) //前台用户API
	}

	merchantPrivateRouter := r.Group("merchant")
	merchantPrivateRouter.Use(middleware.MerchantJWTAuth())
	{
		a.UserAPI.MerchantUserApi.InitRouter(merchantPrivateRouter) //商家用户API
	}

	return r
}
