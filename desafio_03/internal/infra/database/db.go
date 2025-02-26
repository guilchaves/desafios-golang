package database

import (
	"log/slog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		panic(err)
	}

	slog.Info("Database connection established succesfully")
	return db
}
