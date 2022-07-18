package user

import (
	"github.com/gin-gonic/gin"
	"template/internal/app/service/user"
)

type FrontUserApi struct {
	FrontUserService *user.FrontUserService
}

func (api *FrontUserApi) InitRouter(router *gin.RouterGroup) {
	//可选配中间件
	{
		router.POST("/login", api.Login)
	}

}

func (api *FrontUserApi) Login(c *gin.Context) {
	api.FrontUserService.Login()
}
