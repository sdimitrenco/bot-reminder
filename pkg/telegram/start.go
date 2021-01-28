package telegram

import (
	"context"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/controlers"
	"github.com/yanzay/tbot/v2"
	"os"
)

func Start(ctx context.Context, client *tbot.Client, server *tbot.Server) {

	server.HandleMessage("/start", func(m *tbot.Message) {
		controlers.AddNewUser(
			ctx, m.From.Username,
			m.From.FirstName,
			m.From.LastName,
			m.Chat.ID)

		_ = client.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		_, _ = client.SendMessage(m.Chat.ID, m.From.FirstName, tbot.OptReplyKeyboardMarkup(MenuButtons()))
		_, _ = client.SendMessage(m.Chat.ID, os.Getenv("FIRST_TEXT"), tbot.OptParseModeMarkdown)
	})
}
