package messages

import (
	"context"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/controlers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/repositories"
	"github.com/yanzay/tbot/v2"
	"gorm.io/gorm"
	"time"
)

func LeadingTomorrow(ctx context.Context, client *tbot.Client) {
	db := ctx.Value("db").(*gorm.DB)

	userRepo := repositories.NewUser(db)
	arr := userRepo.GetAllRecords()

	t := time.Now().Add(24 * time.Hour)
	text := "Завтра ведет:\n"
	name := controlers.GetUser(ctx, t)

	if name != "Отдыхаем, никто не ведет" {
		for _, value := range arr {
			chatId := value.ChatId
			_, _ = client.SendMessage(chatId, text+name, tbot.OptParseModeHTML)
		}
	}
}

func LeadingToday(ctx context.Context, client *tbot.Client) {
	db := ctx.Value("db").(*gorm.DB)

	userRepo := repositories.NewUser(db)
	arr := userRepo.GetAllRecords()

	t := time.Now()
	text := "Сегодня ведёт:\n"
	name := controlers.GetUser(ctx, t)

	if name != "Отдыхаем, никто не ведет" {
		for _, value := range arr {
			chatId := value.ChatId
			_, _ = client.SendMessage(chatId, text+name, tbot.OptParseModeHTML)
		}
	}
}
