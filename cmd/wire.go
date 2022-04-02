// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"mi/internal/app"
	"mi/internal/app/api"
	"mi/internal/app/repo"
	"mi/internal/app/router"
	"mi/internal/app/service"
)

func initApp() (*app.App, func(), error) {
	panic(wire.Build(
		repo.ProviderSet,
		service.ProviderSet,
		api.ProviderSet,
		router.ProviderSet,
		app.InitEngine,
		newApp,
	))
}
