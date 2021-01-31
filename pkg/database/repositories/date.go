package repositories

import (
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/models"
	"gorm.io/gorm"
	"time"
)

type Date struct {
	db *gorm.DB
}

func NewDate(db *gorm.DB) *Date {
	return &Date{db: db}
}

func (d Date) Create(date *models.Date) *models.Date {
	d.db.Create(&date)
	return date
}

func (d Date) Save(date *models.Date, t time.Time, name string) *models.Date {
	d.db.Model(date).Where("date = ?", t).Update("people_name", name)
	return date
}

func (d Date) FindByDate(date time.Time) *models.Date {
	var dt models.Date
	d.db.First(&dt, "date = ?", date)
	return &dt
}

func (d Date) GetAll() []models.Date {
	var dates []models.Date
	d.db.Find(&dates)
	return dates
}
