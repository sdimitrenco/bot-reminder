package repositories

import (
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/models"
	"gorm.io/gorm"
)

type DailyText struct {
	db *gorm.DB
}

func NewDailyText(db *gorm.DB) *DailyText {
	return &DailyText{db: db}
}

func (repo DailyText) Create(dt *models.DailyText) {

}
