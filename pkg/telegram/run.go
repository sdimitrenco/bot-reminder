package telegram

import (
	"context"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/cron"
	"github.com/yanzay/tbot/v2"
	"os"
)

//running telegram server bot
func Run(ctx context.Context) {
	token := os.Getenv("TELEGRAM_TOKEN")
	bot := tbot.New(token)
	c := bot.Client()
	Handle(ctx, c, bot)
	cron.SenderAutoMessages(ctx, c)
	_ = bot.Start()
}
