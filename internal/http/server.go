package http

import (
	"context"
	"github.com/gofiber/utils/v2"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Jxckaroo/go-fiber-starter-kit/api/middleware"
	"github.com/Jxckaroo/go-fiber-starter-kit/api/router"
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/Jxckaroo/go-fiber-starter-kit/internal/database"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

func NewInstance(c *config.Config) *fiber.App {
	return fiber.New(
		fiber.Config{
			ServerHeader: c.App.Name,
			AppName:      c.App.Name,
			IdleTimeout:  c.App.IdleTimeout * time.Second,
		},
	)
}

func StartServer(
	lifecycle fx.Lifecycle,
	config *config.Config,
	fiber *fiber.App,
	router *router.Router,
	middleware *middleware.Middleware,
	database *database.Database,
	log zerolog.Logger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				host, port := config.App.Host, config.App.Port

				middleware.Register()
				router.RegisterRoutes()

				// ASCII Art
				ascii, err := os.ReadFile("./storage/ascii_art.txt")
				if err != nil {
					log.Debug().Err(err).Msg("An unknown error occurred when to print ASCII art!")
				}

				for _, line := range strings.Split(utils.UnsafeString(ascii), "\n") {
					log.Info().Msg(line)
				}

				log.Info().Msg(fiber.Config().AppName + " is running at the moment!")

				if config.App.Environment != "production" {
					prefork := "Enabled"
					procs := runtime.GOMAXPROCS(0)
					if !config.App.Prefork {
						procs = 1
						prefork = "Disabled"
					}

					log.Debug().Msgf("Version: %s", "-")
					log.Debug().Msgf("Host: %s", host)
					log.Debug().Msgf("Port: %s", port)
					log.Debug().Msgf("Prefork: %s", prefork)
					log.Debug().Msgf("Handlers: %d", fiber.HandlersCount())
					log.Debug().Msgf("Processes: %d", procs)
					log.Debug().Msgf("PID: %d", os.Getpid())

					// Listen the app (with TLS Support)
					// if config.App.Tls.Enabled {
					// 	log.Debug().Msg("TLS support was enabled.")

					// 	if err := fiber.Listen(config.App.Port, config.App.Tls.Cert, config.App.Tls.Key); err != nil {
					// 		log.Error().Err(err).Msg("An unknown error occurred when to run server!")
					// 	}
					// }

					go func() {
						if fiberErr := fiber.Listen(host + ":" + port); err != nil {
							log.Error().Err(fiberErr).Msg("An unknown error occurred when to run server!")
						}
					}()

					// database.ConnectDatabase()

					// migrate := flag.Bool("migrate", false, "migrate the database")
					// seeder := flag.Bool("seed", false, "seed the database")
					// flag.Parse()

					// // read flag -migrate to migrate the database
					// if *migrate {
					// 	database.MigrateModels()
					// }
					// // read flag -seed to seed the database
					// if *seeder {
					// 	database.SeedModels()
					// }

					return nil
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				log.Info().Msg("Shutting down the app...")
				if err := fiber.Shutdown(); err != nil {
					log.Panic().Err(err).Msg("")
				}

				// log.Info().Msg("Running cleanup tasks...")
				// log.Info().Msg("1- Shutdown the database")
				// database.ShutdownDatabase()
				// log.Info().Msgf("%s was successful shutdown.", cfg.App.Name)
				// log.Info().Msg("\u001b[96msee you again👋\u001b[0m")

				return nil
			},
		},
	)
}
