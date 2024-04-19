package router

import (
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	App    fiber.Router
	Config *config.Config
}

func NewRouter(fiber *fiber.App, config *config.Config) *Router {
	return &Router{
		App:    fiber,
		Config: config,
	}
}
