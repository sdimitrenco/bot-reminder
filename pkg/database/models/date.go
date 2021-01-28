package models

import "time"

type Date struct {
	Id         int32 `gorm:"primaryKey"`
	Date       time.Time
	PeopleName string
}

func (d *Date) GetDate() time.Time {
	return d.Date
}

func (d *Date) GetPeopleName() string {
	return d.PeopleName
}
