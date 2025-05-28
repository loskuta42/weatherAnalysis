package fetcher_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/loskuta42/weatherAnalysis/internal/config"
	"github.com/loskuta42/weatherAnalysis/internal/fetcher"
)

func TestFetchWeatherData(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"now_dt": "2025-05-19T09:00:00Z",
			"info": { "lat": 55.75, "lon": 37.61 },
			"fact": { "temp": 21 }
		}`))
	}))
	defer mockServer.Close()

	config.SetConfig(&config.Config{
	Cities: map[string]string{
		"MOSCOW": mockServer.URL,
		"PARIS":  mockServer.URL,
		"LONDON": mockServer.URL,
		"NOVOSIBIRSK": mockServer.URL,
		"CAIRO": mockServer.URL,
	},
	BadCondition: map[string]string{},
	FieldsEnToRU: map[string]string{},
	Timeout:      1 * time.Second,
	})

	f := &fetcher.Fetcher{
		Client: mockServer.Client(),
	}
	cities := []string{"MOSCOW", "PARIS", "LONDON", "NOVOSIBIRSK", "CAIRO", "UNKNOWN_CITY"}
	results := f.RunWeatherWorkerPool(cities, 3)

	var count int
	for res := range results {
		count++
		if res.Err != nil {
			if res.City == "UNKNOWN_CITY" {
				continue
			} else {
				t.Errorf("unexpected error for city %s: %v", res.City, res.Err)
			}
		}
		if res.Weather.Fact.Temp != 21 {
			t.Errorf("unexpected temperature for city %s: got %d, want 25", res.City, res.Weather.Fact.Temp)
		}
	}
	if count != 6 {
		t.Errorf("expected 6 results, got %d", count)
	}
}