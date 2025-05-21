package config

import "time"

type Config struct {
	BadCondition map[string]string
	Cities map[string]string
	FieldsEnToRU map[string]string
	Timeout time.Duration
}


func GetConfig() Config {
	return Config{
		BadCondition: BadCondition,
		Cities: Cities,
		FieldsEnToRU: FieldsEnToRu,
		Timeout: 5*time.Second,
	}
}