package main

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// i prefer including static files in the binary.
// more secure than copying them in continer as they can't change after the build.
//
//go:embed static/*
var content embed.FS

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	author := os.Getenv("AUTHOR")
	if author == "" {
		logger.Error("required env var not provided: $AUTHOR")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handleIndex(author))

	fsys, err := fs.Sub(content, "static")
	if err != nil {
		logger.Error("couldn't prepare static files: " + err.Error())
		os.Exit(1)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(fsys))))

	mux.HandleFunc("/health/readiness", readinessHandler)
	mux.HandleFunc("/health/liveness", livenessHandler)

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// graceful shutdown mechanism
	go func() {
		logger.Info("starting serving new connections...")
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("HTTP server error", "err", err)
			os.Exit(1)
		}
		logger.Info("stopped serving new connections")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("server shutdown error", "err", err)
	}

	logger.Info("server is down")
}
