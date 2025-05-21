package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/loskuta42/weatherAnalysis/internal/config"
	"github.com/loskuta42/weatherAnalysis/internal/fetcher"
)

func main() {
	cfg := config.GetConfig()

	f := &fetcher.Fetcher{
		Client: &http.Client{Timeout: cfg.Timeout},
	}
	result, err := f.FetchWeatherData("MOSCOW")
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%v", result)
}