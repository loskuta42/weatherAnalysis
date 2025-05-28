package config

import "time"

type Config struct {
	BadCondition map[string]string
	Cities map[string]string
	FieldsEnToRU map[string]string
	Timeout time.Duration
}


var cfg *Config

func init() {
	cfg = &Config{
		BadCondition:  BadCondition,
		Cities:        Cities,
		FieldsEnToRU:  FieldsEnToRu,
		Timeout:       5 * time.Second,
	}
}

func GetConfig() *Config {
	return cfg
}

func SetConfig(c *Config) {
	cfg = c
}