package h1

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/AgentNemo/goradios"
)

func Hi(respW http.ResponseWriter, r *http.Request) {
	stations := goradios.FetchStations(goradios.StationsByCountry, "Singapore")

	radios, err := json.Marshal(stations)
	if err != nil {
		respW.WriteHeader(http.StatusInternalServerError)

		return
	}

	respW.Header().Set("Content-Type", "application/json")

	if _, err := respW.Write(radios); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
