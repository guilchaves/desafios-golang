package entity

import (
	"errors"
	"regexp"
)

type Attendee struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"type:string;not null"`
	Email      string `gorm:"type:string;not null"`
	ActivityID int
	Activity   Activity `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

var (
	ErrAttendeeNameIsRequired  = errors.New("name is required")
	ErrAttendeeEmailIsRequired = errors.New("email is required")
	ErrAttendeeEmailIsInvalid  = errors.New("email is invalid")
)

func NewAttendee(name, email string, activityID int) (*Attendee, error) {
	attendee := &Attendee{
		Name:       name,
		Email:      email,
		ActivityID: activityID,
	}

	if err := attendee.Validate(); err != nil {
		return nil, err
	}

	return attendee, nil
}

func (a *Attendee) Validate() error {
	if a.Name == "" {
		return ErrAttendeeNameIsRequired
	}

	if a.Email == "" {
		return ErrAttendeeEmailIsRequired
	}

	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	if !regex.MatchString(a.Email) {
		return ErrAttendeeEmailIsInvalid
	}

	return nil

}
