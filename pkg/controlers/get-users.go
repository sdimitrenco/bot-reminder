package controlers

import (
	"context"
	"fmt"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/models"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/repositories"
	"gorm.io/gorm"
	"time"
)

func GetAllUsersDate(ctx context.Context) []models.Date {
	db := ctx.Value("db").(*gorm.DB)
	dateRepo := repositories.NewDate(db)

	return dateRepo.GetAll()

}

func GetUser(ctx context.Context, t time.Time) string {
	db := ctx.Value("db").(*gorm.DB)
	dateRepo := repositories.NewDate(db)
	layout := "2006-01-02"

	timeString := t.Format(layout)

	fmt.Println(timeString)

	t, _ = time.Parse(layout, timeString)

	fmt.Println(t)

	userDate := dateRepo.FindByDate(t)
	if userDate.PeopleName != "" {
		return fmt.Sprintf("<b>%s</b>", userDate.PeopleName)
	}

	return "Отдыхаем, никто не ведет"

}
