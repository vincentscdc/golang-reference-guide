package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	httpRouter := chi.NewRouter()

	stationFetcher := &Stations{}
	httpRouter.Get("/hi", Hi(stationFetcher))
	httpServer := &http.Server{
		Addr:    "localhost:8081",
		Handler: httpRouter,
	}

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Print(err)
	}
}
