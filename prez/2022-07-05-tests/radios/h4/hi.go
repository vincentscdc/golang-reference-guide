package h4

import (
	"log"
	"net/http"

	"github.com/goccy/go-json"

	"gitlab.com/AgentNemo/goradios"
)

//go:generate mockgen -source=./hi.go -destination=./mock/fetcher.go -package=mock
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

type Stations struct {
	Stations []goradios.Station
}

func (s *Stations) FetchStations(by goradios.StationsBy, term string) {
	s.Stations = goradios.FetchStations(by, term)
}

func (s *Stations) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Stations)
}
