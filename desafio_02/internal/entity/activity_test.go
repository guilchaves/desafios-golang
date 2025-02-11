package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewActivity(t *testing.T) {
	c := &Category{
		ID: 1,
	}

	a, err := NewActivity("Atividade 1", "Descrição atividade 1", 100.0, c.ID)
	assert.NoError(t, err)
	assert.NotNil(t, a)
	assert.Equal(t, "Atividade 1", a.Name)
	assert.Equal(t, "Descrição atividade 1", a.Description)
	assert.Equal(t, 100.0, a.Price)
	assert.Equal(t, c.ID, a.CategoryID)

	a, err = NewActivity("", "Descrição atividade 1", 1000, c.ID)
	assert.Error(t, err)
	assert.Nil(t, a)

	a, err = NewActivity("Atividade 1", "", 1000, c.ID)
	assert.Error(t, err)
	assert.Nil(t, a)

	a, err = NewActivity("Atividade 1", "Descrição atividade 1", 0, c.ID)
	assert.Error(t, err)
	assert.Nil(t, a)
}
