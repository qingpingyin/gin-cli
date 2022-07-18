// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
	"go.uber.org/zap"
	"template/internal/app"
	v1 "template/internal/app/api/v1"
	"template/internal/app/config"
	"template/internal/app/repo"
	"template/internal/app/router"
	"template/internal/app/service"
	"template/pkg/gormx"
)

func initApp(cfg *config.Config, logger *zap.Logger) (*app.App, func(), error) {
	panic(wire.Build(
		gormx.InitGorm,
		repo.ProviderSet,
		service.ProviderSet,
		v1.ProviderSet,
		router.RouterProviderSet,
		newApp,
	))
}
