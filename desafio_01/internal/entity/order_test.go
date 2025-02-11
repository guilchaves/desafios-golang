package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	o, err := NewOrder(1034, 150.0, 20.0)
	assert.NotNil(t, o)
	assert.Nil(t, err)
    assert.Equal(t, 1034, o.Code)
    assert.Equal(t, 150.0, o.BaseValue)
    assert.Equal(t, 20.0, o.Discount)
    assert.NotEqual(t, 0, o.BaseValue)
}
