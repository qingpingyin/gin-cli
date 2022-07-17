package user

import "github.com/gin-gonic/gin"

type FrontUserApi struct {
}

func (api *FrontUserApi) Login(c *gin.Context) {

	frontUserService.Login()
}
