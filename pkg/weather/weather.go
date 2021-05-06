package weather

import (
	"encoding/xml"
	"github.com/StanislavDimitrenco/bot-reminder/pkg/parsers"
	"log"
)

func structWeatherData() XMLData {

	url := "https://xml.meteoservice.ru/export/gismeteo/point/2892.xml"

	var result XMLData

	if xmlBytes, err := parsers.WeatherData(url); err != nil {
		log.Printf("Failed to get XML: %v", err)
		return result
	} else {
		_ = xml.Unmarshal(xmlBytes, &result)

		return result
	}
}

func CreateWeatherMessage() string {
	data := structWeatherData()

	var text string

	text += "\n\n\nПОГОДА СЕГОДНЯ\n\n"

	forecasts := data.Report.Town.Forecast

	for _, forecast := range forecasts {

		switch forecast.Tod {
		case "0":
			text += "🌃 Ночью \n"
		case "1":
			text += "🌇 Утром \n"
		case "2":
			text += "🏙️ Днем \n"
		case "3":
			text += "🌆 Вечером \n"
		}

		text += "\nОблачность: "

		switch forecast.Phenomena.Cloudiness {
		case "-1":
			text += "🌫️ туман"
		case "0":
			text += "☀️ ясно"
		case "1":
			text += "🌤️ малооблачно"
		case "2":
			text += "🌥️ облачно"
		case "3":
			text += "☁️ пасмурно"
		}

		switch forecast.Phenomena.Precipitation {
		case "3":
			text += ", осадки: 🌦️ смешанные\n"
		case "4":
			text += ", осадки: 🌧️ дождь \n"
		case "5":
			text += ", осадки: 🌤️ ливень \n"
		case "6":
			text += ", осадки: 🌧️️ ливень \n"
		case "7":
			text += ", осадки: 🌧️️ ливень \n"
		case "8":
			text += ", осадки: ⛈️️ гроза \n"
		}

		text += "\nТемпература: "
		text += "🌡️ " + forecast.Temperature.Min + " ~ " + forecast.Temperature.Max + "\n"

		text += "Скорость ветра: "
		text += "🌬️ " + forecast.Wind.Min + " ~ " + forecast.Wind.Max + ", "

		switch forecast.Wind.Direction {
		case "1":
			text += "северный"
		case "2":
			text += "северо-восточный,"
		case "3":
			text += "восточный"
		case "4":
			text += "юго-восточный"
		case "5":
			text += "южный"
		case "6":
			text += "юго-западный"
		case "7":
			text += "западный"
		case "8":
			text += "северо-западный"
		}

		text += "\n\n"

	}

	return text
}
