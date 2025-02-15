package utils

import (
	"testing"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entity.Client{})
	assert.NoError(t, err)

	return db
}
