package config

var Cities = map[string]string{
    "MOSCOW": "https://code.s3.yandex.net/async-module/moscow-response.json",
    "PARIS": "https://code.s3.yandex.net/async-module/paris-response.json",
    "LONDON": "https://code.s3.yandex.net/async-module/london-response.json",
    "BERLIN": "https://code.s3.yandex.net/async-module/berlin-response.json",
    "BEIJING": "https://code.s3.yandex.net/async-module/beijing-response.json",
    "KAZAN": "https://code.s3.yandex.net/async-module/kazan-response.json",
    "SPETERSBURG": "https://code.s3.yandex.net/async-module/spetersburg-response.json",
    "VOLGOGRAD": "https://code.s3.yandex.net/async-module/volgograd-response.json",
    "NOVOSIBIRSK": "https://code.s3.yandex.net/async-module/novosibirsk-response.json",
    "KALININGRAD": "https://code.s3.yandex.net/async-module/kaliningrad-response.json",
    "ABUDHABI": "https://code.s3.yandex.net/async-module/abudhabi-response.json",
    "WARSZAWA": "https://code.s3.yandex.net/async-module/warszawa-response.json",
    "BUCHAREST": "https://code.s3.yandex.net/async-module/bucharest-response.json",
    "ROMA": "https://code.s3.yandex.net/async-module/roma-response.json",
    "CAIRO": "https://code.s3.yandex.net/async-module/cairo-response.json",
}

var FieldsEnToRu = map[string]string{
    "ABUDHABI": "Абу Даби",
    "BEIJING": "Пекин",
    "BERLIN": "Берлин",
    "BUCHAREST": "Бухарест",
    "CAIRO": "Каир",
    "KALININGRAD": "Калининград",
    "KAZAN": "Казань",
    "LONDON": "Лондон",
    "MOSCOW": "Москва",
    "NOVOSIBIRSK": "Новосибирск",
    "PARIS": "Париж",
    "ROMA": "Рим",
    "SPETERSBURG": "Санкт-Петербург",
    "VOLGOGRAD": "Волгоград",
    "WARSZAWA": "Варшава",
    "avg_temp": "Температура, среднее",
    "cond_hours": "Без осадков, часов",
    "AVG": "Среднее",
    "dates": "День",
    "rating": "Рейтинг",
    "city_name": "Город",
}

var BadCondition = map[string]string {
	"clear": "ясно",
	"partly-cloudy": "малооблачно",
	"cloudy": "облачно с прояснениями",
	"overcast": "пасмурно",
	"drizzle": "морось",
	"light-rain": "небольшой дождь",
	"rain": "дождь",
	"moderate-rain": "умеренно сильный дождь",
	"heavy-rain": "сильный дождь",
	"continuous-heavy-rain": "длительный сильный дождь",
	"showers": "ливень",
	"wet-snow": "дождь со снегом",
	"light-snow": "небольшой снег",
	"snow": "снег",
	"snow-showers": "снегопад",
	"hail": "град",
	"thunderstorm": "гроза",
	"thunderstorm-with-rain": "дождь с грозой",
	"thunderstorm-with-hail": "гроза с градом",
}