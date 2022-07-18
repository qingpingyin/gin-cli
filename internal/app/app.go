package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"template/internal/app/config"
	"time"
)

type App struct {
	opts   options
	ctx    context.Context
	cancel func()
	lk     sync.Mutex
}

func New(opts ...Option) *App {
	o := options{
		ctx: context.Background(),
	}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}

func (a *App) Run(cfg *config.Config) error {
	state := 1

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := a.init(a.ctx, cfg)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		fmt.Printf("Receive signal[%s]\n", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}
	cleanFunc()
	fmt.Println("server exit")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}

func (a *App) init(ctx context.Context, cfg *config.Config) (func(), error) {

	if cfg.System.IsPrintConfig {
		fmt.Println("--------------")
		config.PrintWithJSON(cfg)
		fmt.Println("--------------")
	}

	httpServerCleanFunc := a.initHTTPServer(ctx, cfg, a.opts.engine)

	return func() {
		httpServerCleanFunc()
	}, nil
}

func (a *App) initHTTPServer(ctx context.Context, cfg *config.Config, handler http.Handler) func() {
	addr := fmt.Sprintf("%s:%d", cfg.System.Addr, cfg.System.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		fmt.Printf("server name :%s\n", a.opts.name)
		fmt.Printf("server version :%s\n", a.opts.version)
		fmt.Printf("the server is running at %s\n", addr)
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
