package calendar

import (
	"github.com/jinzhu/now"
	"time"
)

func ThisMonth() (month string, year int) {
	t := time.Now()
	return MonthList[t.Month().String()], t.Year()
}

func NextMonth() (month string, year int) {
	t := now.EndOfMonth().Add(time.Hour)
	return MonthList[t.Month().String()], t.Year()
}

var MonthList = map[string]string{
	"January":   "ЯНВАРЬ",
	"February":  "ФЕВРАЛЬ",
	"March":     "МАРТ",
	"April":     "АПРЕЛЬ",
	"May":       "МАЙ",
	"June":      "ИЮНЬ",
	"July":      "ИЮЛЬ",
	"August":    "АВГУСТ",
	"September": "СЕНТЯБРЬ",
	"October":   "ОКТЯБРЬ",
	"November":  "НОЯБРЬ",
	"December":  "ДЕКАБРЬ",
}
