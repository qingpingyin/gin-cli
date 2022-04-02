package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"gin-cli/internal/app/service"
)

var UserAPIProviderSet = wire.NewSet(wire.Struct(new(UserAPI),"*"))

type UserAPI struct {
	UserSrv *service.UserService
}

func (u *UserAPI) Login(c *gin.Context) {
	// 登陆参数验证

	// 登陆逻辑
	u.UserSrv.Login()
}
