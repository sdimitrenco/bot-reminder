package parsers

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/models"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/database/repositories"
	strip "github.com/grokify/html-strip-tags-go"
	"gorm.io/gorm"
	"net/http"
)

func ParserDailyVerse(ctx context.Context, url string, date string) string {
	db := ctx.Value("db").(*gorm.DB)
	dtRepo := repositories.NewDailyText(db)

	res, _ := http.Get(url)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "Сайт не отвечает"
	}

	doc, _ := goquery.NewDocumentFromReader(res.Body)

	linkAll := doc.Find(".todayItems").Find(".pub-es21").Find(".scalableui")
	title, _ := linkAll.Find("header").Find("h2").Html()
	script, _ := linkAll.Find(".themeScrp").Html()
	text, _ := linkAll.Find(".bodyTxt").Find(".sb").Html()
	script = strip.StripTags(script)
	text = strip.StripTags(text)

	dt := dtRepo.FindByDate(date)

	if dt.Date == "" {
		dtRepo.Create(&models.DailyText{
			Date:   date,
			Title:  title,
			Script: script,
			Text:   text,
		})
	}

	return fmt.Sprintf("<b>%s</b>\n %s\n %s\n", title, script, text)

}
