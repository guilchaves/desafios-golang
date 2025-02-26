package entity_test

import (
	"testing"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/stretchr/testify/assert"
)

var defaultBirthdate = time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)

func TestNewClient(t *testing.T) {
	t.Run("valid client", func(t *testing.T) {
		client, err := entity.NewClient(
			"John Doe",
			"12345678901",
			5000.0,
			defaultBirthdate,
			0,
		)

		assert.NotNil(t, client)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", client.Name)
		assert.Equal(t, "12345678901", client.Cpf)
		assert.Equal(t, 5000.0, client.Income)
		assert.Equal(t, defaultBirthdate, client.BirthDate)
		assert.Equal(t, uint(0), client.Children)
	})

	t.Run("empty name", func(t *testing.T) {
		client, err := entity.NewClient("", "12345678900", 5000.0, defaultBirthdate, 2)

		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Equal(t, "name cannot be empty", err.Error())
	})

	t.Run("future birthdate", func(t *testing.T) {
		client, err := entity.NewClient("Jane Doe", "09876543210", 4500.0, time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC), 1)

		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Equal(t, "birthdate cannot be in the future", err.Error())
	})
}
