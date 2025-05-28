package model

import (
	"time"
)

// type CityForecast struct {
// 	City      string        `json:"city"`
// 	Forecasts []DayForecast `json:"forecasts"`
// }

type YandexWeatherAPIResponse struct {
	Now   int       `json:"now"`
	NowDt time.Time `json:"now_dt"`
	Info  struct {
		N      bool    `json:"n"`
		Geoid  int     `json:"geoid"`
		URL    string  `json:"url"`
		Lat    float64 `json:"lat"`
		Lon    float64 `json:"lon"`
		Tzinfo struct {
			Name   string `json:"name"`
			Abbr   string `json:"abbr"`
			Dst    bool   `json:"dst"`
			Offset int    `json:"offset"`
		} `json:"tzinfo"`
		DefPressureMm int    `json:"def_pressure_mm"`
		DefPressurePa int    `json:"def_pressure_pa"`
		Slug          string `json:"slug"`
		Zoom          int    `json:"zoom"`
		Nr            bool   `json:"nr"`
		Ns            bool   `json:"ns"`
		Nsr           bool   `json:"nsr"`
		P             bool   `json:"p"`
		F             bool   `json:"f"`
		H             bool   `json:"_h"`
	} `json:"info"`
	GeoObject struct {
		District struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"district"`
		Locality struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"locality"`
		Province struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"province"`
		Country struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"country"`
	} `json:"geo_object"`
	Yesterday struct {
		Temp int `json:"temp"`
	} `json:"yesterday"`
	Fact struct {
		ObsTime      int     `json:"obs_time"`
		Uptime       int     `json:"uptime"`
		Temp         int     `json:"temp"`
		FeelsLike    int     `json:"feels_like"`
		Icon         string  `json:"icon"`
		Condition    string  `json:"condition"`
		Cloudness 	 float64 `json:"cloudness"`
		PrecType     int     `json:"prec_type"`
		PrecProb     int     `json:"prec_prob"`
		PrecStrength float64     `json:"prec_strength"`
		IsThunder    bool    `json:"is_thunder"`
		WindSpeed    float64 `json:"wind_speed"`
		WindDir      string  `json:"wind_dir"`
		PressureMm   int     `json:"pressure_mm"`
		PressurePa   int     `json:"pressure_pa"`
		Humidity     int     `json:"humidity"`
		Daytime      string  `json:"daytime"`
		Polar        bool    `json:"polar"`
		Season       string  `json:"season"`
		Source       string  `json:"source"`
		AccumPrec    struct {
			Num1 float64 `json:"1"`
			Num3 float64 `json:"3"`
			Num7 float64 `json:"7"`
		} `json:"accum_prec"`
		SoilMoisture float64 `json:"soil_moisture"`
		SoilTemp     int     `json:"soil_temp"`
		UvIndex      int     `json:"uv_index"`
		WindGust     float64 `json:"wind_gust"`
	} `json:"fact"`
	Forecasts []struct {
		Date      string `json:"date"`
		DateTs    int    `json:"date_ts"`
		Week      int    `json:"week"`
		Sunrise   string `json:"sunrise"`
		Sunset    string `json:"sunset"`
		RiseBegin string `json:"rise_begin"`
		SetEnd    string `json:"set_end"`
		MoonCode  int    `json:"moon_code"`
		MoonText  string `json:"moon_text"`
		Hours     []struct {
			Hour         string  `json:"hour"`
			HourTs       int     `json:"hour_ts"`
			Temp         int     `json:"temp"`
			FeelsLike    int     `json:"feels_like"`
			Icon         string  `json:"icon"`
			Condition    string  `json:"condition"`
			Cloudness    float64     `json:"cloudness"`
			PrecType     int     `json:"prec_type"`
			PrecStrength float64     `json:"prec_strength"`
			IsThunder    bool    `json:"is_thunder"`
			WindDir      string  `json:"wind_dir"`
			WindSpeed    float64 `json:"wind_speed"`
			WindGust     float64 `json:"wind_gust"`
			PressureMm   int     `json:"pressure_mm"`
			PressurePa   int     `json:"pressure_pa"`
			Humidity     int     `json:"humidity"`
			UvIndex      int     `json:"uv_index"`
			SoilTemp     int     `json:"soil_temp"`
			SoilMoisture float64 `json:"soil_moisture"`
			PrecMm       float64     `json:"prec_mm"`
			PrecPeriod   int     `json:"prec_period"`
			PrecProb     int     `json:"prec_prob"`
		} `json:"hours"`
		Biomet struct {
			Index     int    `json:"index"`
			Condition string `json:"condition"`
		} `json:"biomet,omitempty"`
	} `json:"forecasts"`
}