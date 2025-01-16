package main

import (
	"fmt"
	"log/slog"

	"github.com/JoseLora/fiberapp/internal/infrastructure/di"
)

func main() {
	slog.Info("Starting App")

	if app, error := di.InitializeApp(); error != nil {
		// Log full stack trace error using slog
		slog.Error(fmt.Errorf("Error initializing app: %w", error).Error())
		return
	} else {
		slog.Error(app.Start().Error())
	}
}
