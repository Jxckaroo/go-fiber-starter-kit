package main

import (
	"slices"

	"github.com/Jxckaroo/go-fiber-starter-kit/api/middleware"
	"github.com/Jxckaroo/go-fiber-starter-kit/api/modules"
	"github.com/Jxckaroo/go-fiber-starter-kit/api/router"
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/database"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/logging"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/server"
	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

func main() {
	bootstrap := []fx.Option{
		fx.Provide(config.New),
		fx.Provide(logging.New),
		fx.Provide(server.New),
		fx.Provide(database.New),
		fx.Provide(middleware.New),
		fx.Provide(router.New),
	}

	apiModules := []fx.Option{}
	for _, m := range modules.ToLoad() {
		apiModules = append(apiModules, m.New())
	}

	run := []fx.Option{
		fx.Invoke(server.Start),
		fx.WithLogger(fxzerolog.Init()),
	}

	fx.New(
		slices.Concat(
			bootstrap,
			apiModules,
			run,
		)...,
	).Run()
}
