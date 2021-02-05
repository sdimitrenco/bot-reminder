package controlers

import (
	"context"
	"fmt"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/repositories"
	"gorm.io/gorm"
	"time"
)

func GetAllUsersDate(ctx context.Context) {
	db := ctx.Value("db").(*gorm.DB)
	dateRepo := repositories.NewDate(db)

	allDate := dateRepo.GetAll()

	for _, value := range allDate {
		fmt.Println(value.PeopleName)
		fmt.Println(value.Date.Format("01-02"))

	}
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
