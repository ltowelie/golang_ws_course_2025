package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"repository_example/internal/initialization"
	"syscall"
)

func main() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	app, err := initialization.NewApplication(ctx)
	if err != nil {
		slog.Error("Error initializing application", "error", err)
		os.Exit(1)
	}
	slog.Info("Application initialized")

	_ = app

	// Graceful shutdown
	<-signalChannel
	slog.Info("Received signal, shutting down")
	app.Close(ctx)
	cancel()
}
