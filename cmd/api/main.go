// Package main is the entry point for the Glofox bookings API.
//
// The skeleton starts an HTTP server on :8080 with no routes registered.
// You decide the package layout, the router, and the storage approach.
package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// SeedClass holds the data for the class that is pre-loaded into the store at
// startup. Candidates do not need to create a class — they just book against
// this one. Seed it into your store before registering routes.
var SeedClass = struct {
	ID        string
	Name      string
	StartDate string
	EndDate   string
	Capacity  int
}{
	ID:        "5d2e2a9f-3c1b-4f2a-b6e0-1a2b3c4d5e6f",
	Name:      "Pilates",
	StartDate: "2026-12-14",
	EndDate:   "2026-12-14",
	Capacity:  10,
}

func main() {
	addr := ":8080"
	if v := os.Getenv("ADDR"); v != "" {
		addr = v
	}

	// TODO: initialise your store and seed SeedClass into it here.
	log.Printf("seeded class id=%s name=%q dates=%s..%s capacity=%d",
		SeedClass.ID, SeedClass.Name, SeedClass.StartDate, SeedClass.EndDate, SeedClass.Capacity)

	mux := http.NewServeMux()

	// TODO: register your routes here.
	//
	//   POST /bookings -> book a member onto a pre-seeded class for a specific date
	//
	// See README.md for the full task brief.

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
