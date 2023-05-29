package entity

import (
	"time"
)

type Day uint8

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type WorkingHour struct {
	ID     int64
	UserID int64
	Day    Day
	Start  string
	End    string
}

func (d Day) String() string {
	switch d {
	case 0:
		return "monday"
	case 1:
		return "tuesday"
	case 2:
		return "wednesday"
	case 3:
		return "thrusday"
	case 4:
		return "friday"
	case 5:
		return "saturday"
	case 6:
		return "sunday"
	}

	return ""
}

func (w WorkingHour) IsOnWorkingTime(start, end time.Time, taregtTz *time.Location) (bool, error) {
	dStart, err := time.Parse(time.TimeOnly, w.Start)
	if err != nil {
		return false, err
	}

	dEnd, err := time.Parse(time.TimeOnly, w.End)
	if err != nil {
		return false, err
	}

	convStart := start.In(taregtTz)
	convEnd := end.In(taregtTz)

	dStart = time.Date(convStart.Year(), convStart.Month(), convStart.Day(), dStart.Hour(), dStart.Minute(), dStart.Second(), 0, convStart.Location())
	dEnd = time.Date(convEnd.Year(), convEnd.Month(), convEnd.Day(), dEnd.Hour(), dEnd.Minute(), dEnd.Second(), 0, convEnd.Location())

	validStart := convStart.After(dStart) || convStart.Equal(dStart)
	validEnd := convEnd.Before(dEnd) || convEnd.Equal(dEnd)

	return validStart && validEnd, nil
}
