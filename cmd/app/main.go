package main

import (
	"github.com/Jxckaroo/go-fiber-starter-kit/api/middleware"
	"github.com/Jxckaroo/go-fiber-starter-kit/api/router"
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/database"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/infrastructure"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/logging"
	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
	"slices"
)

func main() {
	bootstrap := []fx.Option{
		fx.Provide(config.New),
		fx.Provide(logging.New),
		fx.Provide(infrastructure.New),
		fx.Provide(database.New),
		fx.Provide(middleware.New),
		fx.Provide(router.New),
	}

	modules := []fx.Option{}

	run := []fx.Option{
		fx.Invoke(infrastructure.Start),
		fx.WithLogger(fxzerolog.Init()),
	}

	fx.New(slices.Concat(bootstrap, modules, run)...).Run()
}
