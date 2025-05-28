package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/loskuta42/weatherAnalysis/internal/config"
	"github.com/loskuta42/weatherAnalysis/internal/model"
)

type Fetcher struct {
	Client *http.Client
}

type WeatherResult struct {
	City string
	Weather model.YandexWeatherAPIResponse
	Err error 
}

func (f *Fetcher) fetchWeatherData(city string) (WeatherResult) {
	cfg := config.GetConfig()
	var weather model.YandexWeatherAPIResponse
	url, exists := cfg.Cities[city]

	if !exists {
		return WeatherResult{City: city, Weather: weather, Err: fmt.Errorf("unknown city: %s", city)}
	}

	resp, err := f.Client.Get(url)
	if err != nil {
		return WeatherResult{City: city, Weather: weather, Err: fmt.Errorf("query error: %v", err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherResult{City: city, Weather: weather, Err: fmt.Errorf("API returned status %s (%d)", resp.Status, resp.StatusCode)}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return WeatherResult{City: city, Weather: weather, Err: fmt.Errorf("failed to read response body: %w", err)}
	}
	
	if err := json.Unmarshal(body, &weather); err != nil {
		return WeatherResult{City: city, Weather: weather, Err: fmt.Errorf("failed to parse weather data: %w", err)} 
	}
	return WeatherResult{City: city, Weather: weather, Err: nil}

}


func (f *Fetcher) RunWeatherWorkerPool(cities []string, numWorkers int) <-chan WeatherResult {
    jobs := make(chan string)
    results := make(chan WeatherResult)
    var wg sync.WaitGroup

    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for city := range jobs {
                res := f.fetchWeatherData(city)
                results <- res
            }
        }()
    }

    go func() {
        for _, city := range cities {
            jobs <- city
        }
        close(jobs)
    }()

    go func() {
        wg.Wait()
        close(results)
    }()

    return results
}
	
