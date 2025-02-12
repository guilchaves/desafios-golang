package entity

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeBlock(t *testing.T) {
	a := &Activity{
		ID: 1,
	}
	tb, err := NewTimeBlock(
		time.Date(2025, 10, 5, 9, 0, 0, 0, time.UTC),
		time.Date(2025, 10, 5, 12, 0, 0, 0, time.UTC),
		a.ID,
	)
	assert.NoError(t, err)
	assert.NotNil(t, tb)

	tb, err = NewTimeBlock(time.Time{}, time.Now(), a.ID)
	assert.Error(t, err)
	assert.Nil(t, tb)

	tb, err = NewTimeBlock(time.Now(), time.Time{}, a.ID)
	assert.Error(t, err)
	assert.Nil(t, tb)

	tb, err = NewTimeBlock(time.Now(), time.Now(), 0)
	assert.NoError(t, err)
	assert.NotNil(t, tb)

	tb, err = NewTimeBlock(time.Time{}, time.Time{}, 0)
	assert.Error(t, err)
	assert.Nil(t, tb)
}
