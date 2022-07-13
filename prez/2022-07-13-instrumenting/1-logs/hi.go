package main

import (
	"log"
	"net/http"

	json "github.com/goccy/go-json"

	"gitlab.com/AgentNemo/goradios"
)

type stationsFetcherMock struct {
	Stations []goradios.Station
}

func (sfm *stationsFetcherMock) FetchStations(by goradios.StationsBy, term string) {
	sfm.Stations = []goradios.Station{
		{
			Name: "Radio 1",
			URL:  "http://radio1.com",
		},
		{
			Name: "Radio 2",
			URL:  "http://radio2.com",
		},
	}
}

type Stations struct {
	Stations []goradios.Station
}

func (s *Stations) FetchStations(by goradios.StationsBy, term string) {
	s.Stations = goradios.FetchStations(by, term)
}

func (s *Stations) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Stations)
}

type StationsFetcher interface {
	FetchStations(by goradios.StationsBy, term string)
	MarshalJSON() ([]byte, error)
}

func Hi(sf StationsFetcher) http.HandlerFunc {
	return http.HandlerFunc(func(respW http.ResponseWriter, r *http.Request) {
		sf.FetchStations(goradios.StationsByCountry, "Singapore")

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
