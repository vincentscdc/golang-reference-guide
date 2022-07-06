package h2

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/AgentNemo/goradios"
)

type StationsFetcher interface {
	FetchStations(by goradios.StationsBy, term string) []goradios.Station
}

func Hi(sf StationsFetcher) http.HandlerFunc {
	return http.HandlerFunc(func(respW http.ResponseWriter, r *http.Request) {
		stations := sf.FetchStations(goradios.StationsByCountry, "Singapore")

		radios, err := json.Marshal(stations)
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
