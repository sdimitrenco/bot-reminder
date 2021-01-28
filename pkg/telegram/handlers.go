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

	//crn := cron.New()
	//_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 28 22 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	//crn.Start()

	//send today's daily text
	server.HandleMessage("📗 Стих на cегодня", func(m *tbot.Message) {
		db := ctx.Value("db").(*gorm.DB)
		dtRepo := repositories.NewDailyText(db)
		t := time.Now()
		date := t.Format("2006/01/02")
		dt := dtRepo.FindByDate(date)

		text := fmt.Sprintf("🗓️ <b>%s</b>\n\n<i>%s</i> \n\n%s ", dt.Title, dt.Script, dt.Text)
		_, _ = client.SendMessage(m.Chat.ID, text, tbot.OptParseModeHTML)
	})

	//send tomorrow's daily text
	server.HandleMessage("📘 Стих на завтра", func(m *tbot.Message) {
		db := ctx.Value("db").(*gorm.DB)
		dtRepo := repositories.NewDailyText(db)
		t := time.Now().Add(24 * time.Hour)
		date := t.Format("2006/01/02")
		dt := dtRepo.FindByDate(date)

		text := fmt.Sprintf("🗓️ <b>%s</b>\n\n<i>%s</i> \n\n%s ", dt.Title, dt.Script, dt.Text)
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

}
