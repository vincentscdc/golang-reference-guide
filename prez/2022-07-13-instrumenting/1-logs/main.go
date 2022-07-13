package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	stationFetcher := &Stations{}

	httpRouter := chi.NewRouter()
	httpRouter.Use(middleware.Logger)
	httpRouter.Get("/hi", Hi(stationFetcher))

	port := 8081
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: httpRouter,
	}

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Print(err)
	}
}
