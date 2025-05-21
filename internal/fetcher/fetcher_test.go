package fetcher_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

	cfg := config.GetConfig()
	cfg.Cities["MOSCOW"] = mockServer.URL

	f := &fetcher.Fetcher{
		Client: mockServer.Client(),
	}

	resp, err := f.FetchWeatherData("MOSCOW")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Fact.Temp != 21 {
		t.Errorf("expected temp 21, got %d", resp.Fact.Temp)
	}
}