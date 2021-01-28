package models

type DailyText struct {
	Date   string
	Title  string
	Script string
	Text   string
}

func (dt *DailyText) GetDate() string {
	return dt.Date
}

func (dt *DailyText) GetTitle() string {
	return dt.Title
}

func (dt *DailyText) GetScript() string {
	return dt.Script
}

func (dt *DailyText) GetText() string {
	return dt.Text
}
