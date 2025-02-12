package entity

import (
	"errors"
	"time"
)

type TimeBlock struct {
	ID         int       `gorm:"primaryKey"`
	StartTime  time.Time `gorm:"type:timestamp;not null"`
	EndTime    time.Time `gorm:"type:timestamp;not null"`
	ActivityID int       
	Activity   Activity  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

var (
	ErrTimeBlockStartTimeIsRequired = errors.New("start time is required")
	ErrTimeBlockEndTimeIsRequired   = errors.New("end time is required")
	ErrTimeBlockActivityIsRequired  = errors.New("activity is required")
	ErrTimeBlockInvalidDuration     = errors.New("end time must be after start time")
	ErrTimeBlockStartIsRequired     = errors.New("start time is required")
	ErrTimeBlockEndIsRequired       = errors.New("end time is required")
	ErrTimeBlockInvalidRange        = errors.New("end time must be after start time")
)

func NewTimeBlock(startTime, endTime time.Time, activityID int) (*TimeBlock, error) {
	timeBlock := &TimeBlock{
		StartTime:  startTime,
		EndTime:    endTime,
		ActivityID: activityID,
	}
	if err := timeBlock.Validate(); err != nil {
		return nil, err
	}
	return timeBlock, nil
}

func (tb *TimeBlock) Validate() error {
	if tb.StartTime.IsZero() {
		return ErrTimeBlockStartTimeIsRequired
	}
	if tb.EndTime.IsZero() {
		return ErrTimeBlockEndTimeIsRequired
	}
	if !tb.EndTime.After(tb.StartTime) {
		return ErrTimeBlockInvalidDuration
	}
	return nil
}
