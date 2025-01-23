package main

import (
	"fmt"
	"log/slog"
	"net/http"
	_ "net/http/pprof"

	"github.com/JoseLora/fiberapp/internal/infrastructure/di"
)

func main() {
	// Start the pprof server in a separate goroutine
	go func() {
		slog.Info("Starting pprof server on :6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			slog.Error("Error starting pprof server: %v", err)
		}
	}()

	slog.Info("Starting App")

	if app, error := di.InitializeApp(); error != nil {
		// Log full stack trace error using slog
		slog.Error(fmt.Errorf("Error initializing app: %w", error).Error())
		return
	} else {
		slog.Error(app.Start().Error())
	}
}
