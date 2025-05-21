package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/loskuta42/weatherAnalysis/internal/config"
	"github.com/loskuta42/weatherAnalysis/internal/model"
)

type Fetcher struct {
	Client *http.Client
}

func (f *Fetcher) FetchWeatherData(city string) (weather model.YandexWeatherAPIResponse, err error) {
	cfg := config.GetConfig()
	
	url, exists := cfg.Cities[city]

	if !exists {
		return weather, fmt.Errorf("unknown city: %s", city)
	}

	resp, err := f.Client.Get(url)
	if err != nil {
		return weather, fmt.Errorf("query error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return weather, fmt.Errorf("API returned status %s (%d)", resp.Status, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weather, fmt.Errorf("failed to read response body: %w", err)
	}
	
	if err := json.Unmarshal(body, &weather); err != nil {
		return weather, fmt.Errorf("failed to parse weather data: %w", err)
	}
	return weather, nil

}