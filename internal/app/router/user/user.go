package user

import (
	"github.com/gin-gonic/gin"
	v1 "template/internal/app/api/v1"
)

type (
	FrontUserRouter    struct{}
	MerchantUserRouter struct{}
)

func (r *FrontUserRouter) InitFrontApiRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")

	userRouterApi := v1.FrontApiGroupApp.UserApiGroup.FrontUserApi
	{
		userRouter.POST("/login", userRouterApi.Login)
	}

}

func (r *MerchantUserRouter) InitMerchantApiRouter(router *gin.RouterGroup) {
	userRouter := router.Group("user")

	userRouterApi := v1.MerchantApiGroupApp.UserApiGroup.MerchantUserApi
	{
		userRouter.POST("/login", userRouterApi.Login)
	}

}
