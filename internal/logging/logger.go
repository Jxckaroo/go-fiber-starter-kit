package logging

import (
	"os"

	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(cfg *config.Config) zerolog.Logger {
	zerolog.TimeFieldFormat = cfg.Logger.TimeFormat

	if cfg.Logger.Prettier {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	zerolog.SetGlobalLevel(zerolog.Level(cfg.Logger.Level))

	return log.Hook(PreforkHook{})
}

type PreforkHook struct{}

func (h PreforkHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if fiber.IsChild() {
		e.Discard()
	}
}
