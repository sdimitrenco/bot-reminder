package cmd

import (
	"context"
	"log"

	"github.com/StanislavDimitrenco/bot-reminder/pkg/calendar"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/controlers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/parsers"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/telegram"
	"github.com/joho/godotenv"
)

//Run - start bot some body
func Run() {
	var ctx context.Context
	//load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading env file")
	}

	//add database to context
	ctx = database.Boot(context.Background())

	nowMonth, nowYear := calendar.ThisMonth()
	parsers.RunTableParser(ctx, nowMonth, nowYear)

	nextMonth, nextYear := calendar.NextMonth()
	parsers.RunTableParser(ctx, nextMonth, nextYear)

	//parser daily text
	go controlers.GetDailyText(ctx)

	telegram.CallBack(ctx)

	//telegram server
	go telegram.Run(ctx)

}
