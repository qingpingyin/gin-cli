package app

import (
	"context"
	"fmt"
	"gin-cli/internal/app/config"
	"log"
	"net/http"
	"time"

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
func SetContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}
func SetName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}
func Init(ctx context.Context, cfg *config.Config, engine *gin.Engine) (func(), error) {

	fmt.Println("--------------")
	config.PrintWithJSON(cfg)
	fmt.Println("--------------")

	httpServerCleanFunc := InitHTTPServer(ctx, cfg, engine)

	return func() {
		httpServerCleanFunc()
	}, nil
}

func InitHTTPServer(ctx context.Context, cfg *config.Config, handler http.Handler) func() {
	addr := fmt.Sprintf("%s:%d", cfg.System.Addr, cfg.System.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		log.Printf("the server is running at %s ", addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(30))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		}
	}
}
