package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/monacohq/golang-common/monitoring/otelinit"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	stationFetcher := &Stations{}

	ctx := context.Background()
	shutdownOtel, err := otelinit.InitProvider(
		ctx,
		"radios",
		otelinit.WithGRPCTraceExporter(
			ctx,
			fmt.Sprintf("%s:%d", "127.0.0.1", 4317),
		),
	)
	if err != nil {
		log.Fatalf("failed to initialize opentelemetry: %v", err)
	}
	defer shutdownOtel()

	httpRouter := chi.NewRouter()
	httpRouter.Use(middleware.Logger)
	httpRouter.Mount("/debug", middleware.Profiler())
	httpRouter.Get("/hi", Hi(stationFetcher))
	httpRouter.Get("/leaky", Leaky())
	httpRouter.Get("/nonleaky", NonLeaky())
	httpRouter.Get("/nonleakyoptim", NonLeakyOptim())

	port := 8081
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", port),
		Handler: otelhttp.NewHandler(httpRouter, "server"),
	}

	if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Print(err)
	}
}
