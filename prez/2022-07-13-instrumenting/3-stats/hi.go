package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	json "github.com/goccy/go-json"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"gitlab.com/AgentNemo/goradios"
)

type Station struct {
	ChangeUUID         string `json:"changeuuid"`
	StationUUID        string `json:"stationuuid"`
	Name               string `json:"name"`
	URL                string `json:"url"`
	URLResolved        string `json:"url_resolved"`
	Homepage           string `json:"homepage"`
	Favicon            string `json:"favicon"`
	Tags               string `json:"tags"`
	Country            string `json:"country"`
	CountryCode        string `json:"countrycode"`
	State              string `json:"state"`
	Language           string `json:"language"`
	Votes              int    `json:"votes"`
	LastChangeTime     string `json:"lastchangetime"`
	Codec              string `json:"codec"`
	Bitrate            int    `json:"bitrate"`
	HLS                int    `json:"hls"`
	LastCheckOk        int    `json:"lastcheckok"`
	LastCheckTime      string `json:"lastchecktime"`
	LastCheckOkTime    string `json:"lastcheckoktime"`
	LastLocalCheckTime string `json:"lastlocalchecktime"`
	ClickTimestamp     string `json:"clicktimestamp"`
	ClickCount         int    `json:"clickcount"`
	ClickTrend         int    `json:"clicktrend"`
}

type stationsFetcherMock struct {
	Stations []*Station
}

func (sfm *stationsFetcherMock) FetchStations(ctx context.Context, by goradios.StationsBy, term string) error {
	sfm.Stations = []*Station{
		{
			Name: "Radio 1",
			URL:  "http://radio1.com",
		},
		{
			Name: "Radio 2",
			URL:  "http://radio2.com",
		},
	}

	return nil
}

type Stations struct {
	Stations []*Station
}

func (s *Stations) FetchStations(ctx context.Context, by goradios.StationsBy, term string) error {
	res, err := otelhttp.Get(ctx, goradios.GenerateStationsURL(by, term))
	if err != nil {
		return fmt.Errorf("FetchStations: %w", err)
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(&s.Stations)
}

func (s *Stations) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Stations)
}

type StationsFetcher interface {
	FetchStations(ctx context.Context, by goradios.StationsBy, term string) error
	MarshalJSON() ([]byte, error)
}

func Hi(sf StationsFetcher) http.HandlerFunc {
	return http.HandlerFunc(func(respW http.ResponseWriter, r *http.Request) {
		if err := sf.FetchStations(r.Context(), goradios.StationsByCountry, "Singapore"); err != nil {
			log.Printf("%v", err)
			respW.WriteHeader(http.StatusInternalServerError)

			return
		}

		radios, err := sf.MarshalJSON()
		if err != nil {
			respW.WriteHeader(http.StatusInternalServerError)

			return
		}

		respW.Header().Set("Content-Type", "application/json")

		if _, err := respW.Write(radios); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})
}
