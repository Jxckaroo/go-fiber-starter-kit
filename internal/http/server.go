package http

import (
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/gofiber/fiber/v3"
	"time"
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
