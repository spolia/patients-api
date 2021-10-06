package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

type server struct {
	*http.Server
}

func newServer(listening string, mux *chi.Mux) *server {
	s := &http.Server{
		Addr:         ":" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

// Start runs ListenAndServe on the http.Server with graceful shutdown
func (srv *server) Start() {
	log.Println("starting API cmd")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("could not listen on %s due to %s", srv.Addr, err.Error())
		}
	}()
	log.Printf("cmd is ready to handle requests %s", srv.Addr)
	srv.gracefulShutdown()
}

func (srv *server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Printf("cmd is shutting down %s", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	srv.SetKeepAlivesEnabled(false)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("could not gracefully shutdown the cmd %s", err.Error())
	}
	log.Printf("cmd stopped")
}
