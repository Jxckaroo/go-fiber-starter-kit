package config

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	App struct {
		Name        string        `env:"APP_NAME"`
		Host        string        `env:"APP_HOST"`
		Port        string        `env:"APP_PORT"`
		IdleTimeout time.Duration `env:"APP_IDLE_TIMEOUT"`
		PrintRoutes bool          `env:"APP_PRINT_ROUTES"`
		Prefork     bool          `env:"APP_PREFORK"`
		Environment string        `env:"APP_ENVIRONMENT"`
		Tls         struct {
			Enabled bool   `env:"APP_TLS_ENABLED"`
			Cert    string `env:"APP_TLS_CERT"`
			Key     string `env:"APP_TLS_KEY"`
		}
	}
	Database struct {
		Dsn string `env:"DATABASE_DSN"`
	}
	Logger struct {
		TimeFormat string `env:"LOGGER_TIME_FORMAT"`
		Level      int8   `env:"LOGGER_LEVEL"`
		Prettier   bool   `env:"LOGGER_PRETTIER"`
	}
	Middleware struct {
		Compress struct {
			Enable bool `env:"MIDDLEWARE_COMPRESS_ENABLE"`
			Level  int  `env:"MIDDLEWARE_COMPRESS_LEVEL"`
		}
		Recover struct {
			Enable bool `env:"MIDDLEWARE_RECOVER_ENABLE"`
		}
		Monitor struct {
			Enable bool   `env:"MIDDLEWARE_MONITOR_ENABLE"`
			Path   string `env:"MIDDLEWARE_MONITOR_PATH"`
		}
		Pprof struct {
			Enable bool `env:"MIDDLEWARE_PPROF_ENABLE"`
		}
		Limiter struct {
			Enable     bool          `env:"MIDDLEWARE_LIMITER_ENABLE"`
			Max        int           `env:"MIDDLEWARE_LIMITER_MAX"`
			Expiration time.Duration `env:"MIDDLEWARE_LIMITER_EXPIRATION"`
		}
		Jwt struct {
			Secret     string        `env:"MIDDLEWARE_JWT_SECRET"`
			Expiration time.Duration `env:"MIDDLEWARE_JWT_EXPIRATION"`
		}
		FileSystem struct {
			Enable bool   `env:"MIDDLEWARE_FILESYSTEM_ENABLE"`
			Browse bool   `env:"MIDDLEWARE_FILESYSTEM_BROWSE"`
			MaxAge int    `env:"MIDDLEWARE_FILESYSTEM_MAX_AGE"`
			Index  string `env:"MIDDLEWARE_FILESYSTEM_INDEX"`
			Root   string `env:"MIDDLEWARE_FILESYSTEM_ROOT"`
		}
	}
}

func New() *Config {
	ctx := context.Background()

	var c Config
	_ = godotenv.Load()
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}
	return &c
}
