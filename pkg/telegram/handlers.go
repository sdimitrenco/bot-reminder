package telegram

import (
	"context"
	"fmt"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/controlers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/repositories"
	"github.com/yanzay/tbot/v2"
	"gorm.io/gorm"
	"time"
)

func Handle(ctx context.Context, client *tbot.Client, server *tbot.Server) {
	Start(ctx, client, server)

	//send today's daily text
	server.HandleMessage("📗 Стих на cегодня", func(m *tbot.Message) {

		db := ctx.Value("db").(*gorm.DB)
		dtRepo := repositories.NewDailyText(db)
		t := time.Now()
		date := t.Format("2006/1/2")
		dt := dtRepo.FindByDate(date)

		text := fmt.Sprintf("🗓️ <b>%s</b>\n\n<i>%s</i> \n\n%s ", dt.Title, dt.Script, dt.Text)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
	})

	//send tomorrow's daily text
	server.HandleMessage("📘 Стих на завтра", func(m *tbot.Message) {
		db := ctx.Value("db").(*gorm.DB)
		dtRepo := repositories.NewDailyText(db)
		t := time.Now().Add(24 * time.Hour)
		date := t.Format("2006/1/2")
		dt := dtRepo.FindByDate(date)

		text := fmt.Sprintf("🗓️ <b>%s</b>\n\n", dt.Title)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
		text = fmt.Sprintf("🗓️ <b>%s</b>\n\n<i>%s</i> ", dt.Script, dt.Text)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
	})

	//send who have daily text

	server.HandleMessage("🙋 ‍Завтра ведет", func(m *tbot.Message) {
		t := time.Now().Add(24 * time.Hour)
		text := controlers.GetUser(ctx, t)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
	})

	server.HandleMessage("🙋 Сегодня ведет", func(m *tbot.Message) {
		t := time.Now()
		text := controlers.GetUser(ctx, t)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
	})

	server.HandleMessage("📅 Список ведущих за месяц", func(m *tbot.Message) {
		_, _ = client.SendMessage(m.Chat.ID, CallBack(ctx), tbot.OptParseModeHTML)
	})

}

func CallBack(ctx context.Context) (users string) {
	allUsers := controlers.GetAllUsersDate(ctx)
	nowMonth := time.Now().Format("Jan")

	for _, value := range allUsers {
		if nowMonth == value.Date.Format("Jan") {
			if time.Now().Format("2006/1/2") == value.Date.Format("2006/1/2") {
				users += fmt.Sprintf("➡️ <b>%s - %s</b> \n", value.Date.Format("01-02"), value.PeopleName)
			} else {
				users += fmt.Sprintf("%s - %s \n", value.Date.Format("01-02"), value.PeopleName)
			}
		}
	}

	return users

}
