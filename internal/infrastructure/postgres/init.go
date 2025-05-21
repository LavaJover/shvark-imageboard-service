package postgres

import (
	"log"

	"github.com/LavaJover/shvark-imageboard-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustInitDB(cfg *config.ImageboardConfig) *gorm.DB {
	dsn := cfg.ImageboardDB.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init db: %v\n", err)
	}

	db.AutoMigrate(&ReviewModel{}, &ImageModel{})

	return db
}