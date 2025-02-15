package entity

import (
	"testing"
	"time"

	"github.com/guilchaves/desafios-golang/desafio_03/internal/entity"
	"github.com/guilchaves/desafios-golang/desafio_03/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c, err := entity.NewClient(
		"John Doe",
		"12312312312",
		5000.0,
		time.Date(1990, 02, 02, 0, 0, 0, 0, time.UTC),
		0,
	)
	assert.NotNil(t, c)
	assert.NoError(t, err)

	c, err = entity.NewClient(
		"",
		"12312312312",
		5000.0,
		time.Date(1990, 02, 02, 0, 0, 0, 0, time.UTC),
		0,
	)
	assert.Nil(t, c)
	assert.Error(t, err, validator.ErrNameIsEmpty)

	c, err = entity.NewClient(
		"John Doe",
		"123123123",
		5000.0,
		time.Date(1990, 02, 02, 0, 0, 0, 0, time.UTC),
		0,
	)
	assert.Nil(t, c)
	assert.Error(t, err, validator.ErrCpfIsInvalid)

	c, err = entity.NewClient(
		"John Doe",
		"12312312312",
		-5000.0,
		time.Date(1990, 02, 02, 0, 0, 0, 0, time.UTC),
		0,
	)
	assert.Nil(t, c)
	assert.Error(t, err, validator.ErrIncomeIsInvalid)

	c, err = entity.NewClient(
		"John Doe",
		"12312312312",
		5000.0,
		time.Date(2026, 02, 02, 0, 0, 0, 0, time.UTC),
		0,
	)
	assert.Nil(t, c)
	assert.Error(t, err, validator.ErrBirthdayIsFuture)

	c, err = entity.NewClient(
		"John Doe",
		"12312312312",
		5000.0,
		time.Date(1990, 02, 02, 0, 0, 0, 0, time.UTC),
		-2,
	)
	assert.Nil(t, c)
	assert.Error(t, err, validator.ErrChildrenIsInvalid)

}
