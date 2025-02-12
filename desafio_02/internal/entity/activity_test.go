package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewActivity(t *testing.T) {
	c := &Category{
		ID: 1,
	}

	startTime := time.Date(2025, 10, 5, 9, 0, 0, 0, time.UTC)
	endTime := time.Date(2025, 10, 5, 12, 0, 0, 0, time.UTC)
	tb := TimeBlock{StartTime: startTime, EndTime: endTime, ActivityID: 1}

	a, err := NewActivity(
		"Atividade 1",
		"Descrição atividade 1",
		100.0,
		c.ID,
		[]TimeBlock{tb},
	)
	assert.NoError(t, err)
	assert.NotNil(t, a)

	a, err = NewActivity("Atividade 1", "Descrição atividade 1", 100.0, c.ID, []TimeBlock{tb})
	assert.NotNil(t, tb)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, "Atividade 1", a.Name)
	assert.Equal(t, "Descrição atividade 1", a.Description)
	assert.Equal(t, 100.0, a.Price)
	assert.Equal(t, c.ID, a.CategoryID)
	assert.Len(t, a.TimeBlocks, 1)

	a, err = NewActivity("", "Descrição atividade 1", 100.0, c.ID, []TimeBlock{tb})
	assert.ErrorIs(t, err, ErrActivityNameIsRequired)
	assert.Nil(t, a)

	a, err = NewActivity("Atividade 1", "", 100.0, c.ID, []TimeBlock{tb})
	assert.ErrorIs(t, err, ErrActivityDescriptionIsRequired)
	assert.Nil(t, a)

	a, err = NewActivity("Atividade 1", "Descrição atividade 1", 0, c.ID, []TimeBlock{tb})
	assert.ErrorIs(t, err, ErrActivityPriceIsRequired)
	assert.Nil(t, a)

	a, err = NewActivity("Atividade 1", "Descrição atividade 1", -10, c.ID, []TimeBlock{tb})
	assert.ErrorIs(t, err, ErrActivityPriceIsInvalid)
	assert.Nil(t, a)

	a, err = NewActivity("Atividade 1", "Descrição atividade 1", 100.0, 0, []TimeBlock{tb})
	assert.ErrorIs(t, err, ErrActivityCategoryIsRequired)
	assert.Nil(t, a)

}
