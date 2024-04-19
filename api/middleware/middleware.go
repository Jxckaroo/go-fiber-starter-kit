package middleware

import (
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/filesystem"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/pprof"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"os"
	"time"
)

type Middleware struct {
	App    *fiber.App
	Config *config.Config
}

func NewMiddleware(app *fiber.App, config *config.Config) *Middleware {
	return &Middleware{
		App:    app,
		Config: config,
	}
}

// Register registers all the middleware functions
func (m *Middleware) Register() {
	// Add Extra Middlewares
	m.App.Use(limiter.New(limiter.Config{
		Next:       IsEnabled(m.Config.Middleware.Limiter.Enable),
		Max:        m.Config.Middleware.Limiter.Max,
		Expiration: m.Config.Middleware.Limiter.ExpirationSeconds * time.Second,
	}))

	m.App.Use(compress.New(compress.Config{
		Next:  IsEnabled(m.Config.Middleware.Compress.Enable),
		Level: compress.Level(m.Config.Middleware.Compress.Level),
	}))

	m.App.Use(recover.New(recover.Config{
		Next: IsEnabled(m.Config.Middleware.Recover.Enable),
	}))

	m.App.Use(pprof.New(pprof.Config{
		Next: IsEnabled(m.Config.Middleware.Pprof.Enable),
	}))

	m.App.Use(filesystem.New(filesystem.Config{
		Next:   IsEnabled(m.Config.Middleware.FileSystem.Enable),
		Root:   os.DirFS(m.Config.Middleware.FileSystem.Root),
		Browse: m.Config.Middleware.FileSystem.Browse,
		MaxAge: m.Config.Middleware.FileSystem.MaxAge,
	}))
}

func IsEnabled(key bool) func(c fiber.Ctx) bool {
	if key {
		return nil
	}

	return func(c fiber.Ctx) bool {
		return true
	}
}
