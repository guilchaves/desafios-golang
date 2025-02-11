package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCategory(t *testing.T) {
	c, err := NewCategory(1, Oficina)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, 1, c.ID)
	assert.Equal(t, Oficina, c.Description)
	assert.NotNil(t, c.Activities)

	c, err = NewCategory(0, Oficina)
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = NewCategory(2, "")
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = NewCategory(3, "invalid")
	assert.Error(t, err)
	assert.Nil(t, c)
}
