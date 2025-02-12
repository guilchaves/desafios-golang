package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAttendee(t *testing.T) {
	activity := &Activity{
		ID: 1,
	}

	at, err := NewAttendee("John Doe", "john@email.com", activity.ID)
	assert.NotNil(t, at)
	assert.NoError(t, err)
	assert.Equal(t, at.Name, "John Doe")
	assert.Equal(t, at.Email, "john@email.com")
	assert.Equal(t, at.ActivityID, activity.ID)

	at, err = NewAttendee("", "john@email.com", activity.ID)
	assert.Error(t, err)
	assert.Nil(t, at)

	at, err = NewAttendee("John Doe", "", activity.ID)
	assert.Error(t, err)
	assert.Nil(t, at)

	at, err = NewAttendee("John Doe", "john", activity.ID)
	assert.Error(t, err)
	assert.Nil(t, at)

	
}
