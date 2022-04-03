package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"gin-cli/internal/app/api"
)

var ProviderSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

var _ IRouter = (*Router)(nil)

type IRouter interface {
	Register(app *gin.Engine) error
}

type Router struct {
	// CasbinEnforcer *casbin.SyncedEnforcer
	UserAPI *api.UserAPI
}

func (a *Router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *Router) RegisterAPI(app *gin.Engine) {

	g := app.Group("/api")

	// 配置路由中间件
	v1 := g.Group("/v1")
	{
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "ok",
				"status": 200,
			})
		})

		user := v1.Group("/user")
		{
			user.POST("/login", a.UserAPI.Login)
		}

	}

}
