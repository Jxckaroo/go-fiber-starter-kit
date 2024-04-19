package database

import (
	"github.com/Jxckaroo/go-fiber-starter-kit/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Database struct {
	DB     *gorm.DB
	Logger zerolog.Logger
	Config *config.Config
}

func NewDatabase(cfg *config.Config, logger zerolog.Logger) *Database {
	db := &Database{
		Config: cfg,
		Logger: logger,
	}

	return db
}
