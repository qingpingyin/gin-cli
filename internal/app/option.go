package app

import (
	"context"
	"github.com/gin-gonic/gin"
)

type options struct {
	version string
	name    string
	engine  *gin.Engine
	ctx     context.Context
}

type Option func(*options)

func SetEngine(engine *gin.Engine) Option {
	return func(o *options) {
		o.engine = engine
	}
}

func SetVersion(version string) Option {
	return func(o *options) {
		o.version = version
	}
}
func SetName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}
