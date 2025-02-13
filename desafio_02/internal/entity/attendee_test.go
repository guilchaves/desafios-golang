package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttendee(t *testing.T) {
	at, err := NewAttendee("John Doe", "john@email.com")
	assert.NotNil(t, at)
	assert.NoError(t, err)
	assert.Equal(t, at.Name, "John Doe")
	assert.Equal(t, at.Email, "john@email.com")

	at, err = NewAttendee("", "john@email.com")
	assert.Error(t, err)
	assert.Nil(t, at)

	at, err = NewAttendee("John Doe", "")
	assert.Error(t, err)
	assert.Nil(t, at)

	at, err = NewAttendee("John Doe", "john")
	assert.Error(t, err)
	assert.Nil(t, at)

	
}
