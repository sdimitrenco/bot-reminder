package cron

import (
	"context"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/calendar"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/controlers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/parsers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/telegram/messages"
	"github.com/robfig/cron/v3"
	"github.com/yanzay/tbot/v2"
)

func SenderAutoMessages(ctx context.Context, client *tbot.Client) {

	crn := cron.New()
	//name leading tomorrow
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 20 * * *", func() {
		messages.LeadingTomorrow(ctx, client)
	})
	//name leading today
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 09 * * *", func() {
		messages.LeadingToday(ctx, client)
	})
	//parse list users
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 02 02 * * *", func() {
		nowMonth, nowYear := calendar.ThisMonth()
		go parsers.RunTableParser(ctx, nowMonth, nowYear)
	})
	//parse list users next month
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 02 28 * *", func() {
		nextMonth, nextYear := calendar.NextMonth()
		go parsers.RunTableParser(ctx, nextMonth, nextYear)
	})
	//get daily text
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 02 * * *", func() {
		controlers.GetDailyText(ctx)
	})

	crn.Start()
}
