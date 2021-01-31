package cron

import (
	"context"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/calendar"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/parsers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/telegram/messages"
	"github.com/robfig/cron/v3"
	"github.com/yanzay/tbot/v2"
)

func SenderAutoMessages(ctx context.Context, client *tbot.Client) {
	crn := cron.New()
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 20 * * *", func() {
		messages.LeadingTomorrow(ctx, client)
	})
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 08 * * *", func() {
		messages.LeadingToday(ctx, client)
	})
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 02 * * *", func() {
		//parse list users
		nowMonth, nowYear := calendar.ThisMonth()
		go parsers.RunTableParser(ctx, nowMonth, nowYear)
	})
	_, _ = crn.AddFunc("CRON_TZ=Europe/Moscow 00 02 28 * *", func() {
		//parse list users next month
		nextMonth, nextYear := calendar.NextMonth()
		go parsers.RunTableParser(ctx, nextMonth, nextYear)
	})

	crn.Start()
}
