// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/JoseLora/fiberapp/internal/amiga/config"
	"github.com/JoseLora/fiberapp/internal/api/handler"
	"github.com/JoseLora/fiberapp/internal/application/usecase"
	"github.com/JoseLora/fiberapp/internal/infrastructure/repository"
	"github.com/JoseLora/fiberapp/internal/infrastructure/server/eventbus"
	"github.com/JoseLora/fiberapp/internal/infrastructure/server/http"
)

// Injectors from wire.go:

func InitializeApp() (*http.Server, error) {
	product := repository.NewProductInMemory()
	bus := eventbus.NewEventBus()
	configConfig, err := config.NewConfig(bus)
	if err != nil {
		return nil, err
	}
	productFinderAll := usecase.NewProductFinderAll(product, configConfig, bus)
	productFinderByID := usecase.NewProductFinderByID(product)
	productHandler := handler.NewProductAPI(productFinderAll, productFinderByID)
	server := http.NewServer(productHandler)
	return server, nil
}
