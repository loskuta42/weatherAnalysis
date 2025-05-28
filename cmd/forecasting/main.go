package main

import (
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
	results := f.RunWeatherWorkerPool([]string{"MOSCOW"}, 1)
	for res := range results {
		if res.Err != nil { 
			log.Printf("unexpected error for city %s: %v", res.City, res.Err)
			continue
		}
		log.Println(res.Weather)
	}
}