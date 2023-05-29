package workinghourrepo

import (
	"appointment-api/internal/entity"
)

type WorkingHour struct {
	ID     int64 `gorm:"primaryKey"`
	UserID int64
	Day    uint8
	Start  string
	End    string
}

func (WorkingHour) TableName() string {
	return "working_hours"
}

func (w WorkingHour) ToEntity() entity.WorkingHour {
	return entity.WorkingHour{
		ID:     w.ID,
		UserID: w.UserID,
		Day:    entity.Day(w.Day),
		Start:  w.Start,
		End:    w.End,
	}
}
