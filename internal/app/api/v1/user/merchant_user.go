package user

import (
	"github.com/gin-gonic/gin"
	"template/internal/app/service/user"
)

type MerchantUserApi struct {
	UserService *user.MerchantUserService
}

func (m *MerchantUserApi) InitRouter(router *gin.RouterGroup) {
	//可选配中间件
	{
		router.POST("/login", m.Login)
	}
}

func (mu *MerchantUserApi) Login(c *gin.Context) {

}
