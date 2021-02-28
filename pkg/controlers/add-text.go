package controlers

import (
	"context"
	"fmt"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/parsers"
	"os"
	"time"
)

func GetDailyText(ctx context.Context) {

	t := time.Now().Add(24 * time.Hour)
	frt := t.Format("2006/1/2")
	fmt.Println(frt)
	url := os.Getenv("JW_URL") + frt
	parsers.ParserDailyVerse(ctx, url, frt)

}
