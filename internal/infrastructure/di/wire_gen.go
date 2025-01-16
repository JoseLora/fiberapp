// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/JoseLora/fiberapp/internal/api/handler"
	"github.com/JoseLora/fiberapp/internal/application/usecase"
	"github.com/JoseLora/fiberapp/internal/infrastructure/repository"
	"github.com/JoseLora/fiberapp/internal/infrastructure/server/http"
)

// Injectors from wire.go:

func InitializeApp() (*http.Server, error) {
	product := repository.NewProductInMemory()
	productFinderAll := usecase.NewProductFinderAll(product)
	productFinderByID := usecase.NewProductFinderByID(product)
	productAPI := handler.NewProductAPI(productFinderAll, productFinderByID)
	server := http.NewServer(productAPI)
	return server, nil
}
