package main

import (
	"github.com/Jxckaroo/go-fiber-starter-kit/api/middleware"
	"github.com/Jxckaroo/go-fiber-starter-kit/api/router"
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/database"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/http"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/logging"
	fxzerolog "github.com/efectn/fx-zerolog"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.NewConfig),

		fx.Provide(logging.NewLogger),

		fx.Provide(http.NewInstance),

		fx.Provide(database.NewDatabase),

		fx.Provide(middleware.NewMiddleware),

		fx.Provide(router.NewRouter),

		// modules load here

		fx.Invoke(http.StartServer),

		fx.WithLogger(fxzerolog.Init()),
	).Run()
}
