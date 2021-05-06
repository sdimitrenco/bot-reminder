package weather

import "encoding/xml"

type XMLData struct {
	XMLName xml.Name `xml:"MMWEATHER"`
	Report  Report   `xml:"REPORT"`
}

type Report struct {
	Town Town `xml:"TOWN"`
}

type Town struct {
	Forecast []Forecast `xml:"FORECAST"`
}

type Forecast struct {
	Day         string      `xml:"day,attr"`
	Month       string      `xml:"month,attr"`
	Hour        string      `xml:"hour,attr"`
	Tod         string      `xml:"tod,attr"` //время суток, для которого составлен прогноз: 0 - ночь 1 - утро, 2 - день, 3 - вечер
	Phenomena   Phenomena   `xml:"PHENOMENA"`
	Temperature Temperature `xml:"TEMPERATURE"`
	Wind        Wind        `xml:"WIND"`
}

type Phenomena struct {
	Cloudiness    string `xml:"cloudiness,attr"`    //облачность по градациям: -1 - туман, 0 - ясно, 1 - малооблачно, 2 - облачно, 3 - пасмурно
	Precipitation string `xml:"precipitation,attr"` //тип осадков: 3 - смешанные, 4 - дождь, 5 - ливень, 6,7 – снег, 8 - гроза, 9 - нет данных, 10 - без осадков
}

type Temperature struct {
	Max string `xml:"max,attr"` //температура воздуха, в градусах Цельсия
	Min string `xml:"min,attr"` //температура воздуха, в градусах Цельсия
}

type Wind struct {
	Max       string `xml:"max,attr"` //минимальное и максимальное значения средней скорости ветра, без порывов (м/с)
	Min       string `xml:"min,attr"`
	Direction string `xml:"direction,attr"` //направление ветра в румбах от 1 до 8: 1 - северный, 2 - северо-восточный, и т.д.
}
