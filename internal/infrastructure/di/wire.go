//go:build wireinject
// +build wireinject

package di

import (
	amiga "github.com/JoseLora/fiberapp/internal/amiga/config"
	"github.com/JoseLora/fiberapp/internal/api/handler"
	"github.com/JoseLora/fiberapp/internal/application/usecase"
	"github.com/JoseLora/fiberapp/internal/infrastructure/repository"
	"github.com/JoseLora/fiberapp/internal/infrastructure/server/http"
	"github.com/google/wire"
)

func InitializeApp() (*http.Server, error) {
	wire.Build(
		amiga.NewConfig,
		repository.NewProductInMemory,
		handler.NewProductAPI,
		usecase.NewProductFinderAll,
		usecase.NewProductFinderByID,
		http.NewServer,
	)
	return &http.Server{}, nil
}
