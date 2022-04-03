// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"gin-cli/internal/app"
	"gin-cli/internal/app/api"
	"gin-cli/internal/app/config"
	"gin-cli/internal/app/repo"
	"gin-cli/internal/app/router"
	"gin-cli/internal/app/service"
)

func initApp(cfg *config.Config) (*app.App, func(), error) {
	panic(wire.Build(
		app.InitGorm,
		repo.ProviderSet,
		service.ProviderSet,
		api.ProviderSet,
		router.ProviderSet,
		app.InitEngine,
		newApp,
	))
}
