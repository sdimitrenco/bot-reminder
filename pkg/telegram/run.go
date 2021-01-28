package telegram

import (
	"context"
	"github.com/yanzay/tbot/v2"
	"os"
)

//running telegram server bot
func Run(ctx context.Context) {
	token := os.Getenv("TELEGRAM_TOKEN")
	bot := tbot.New(token)
	c := bot.Client()
	Handle(ctx, c, bot)
	_ = bot.Start()
}
