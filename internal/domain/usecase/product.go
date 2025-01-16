package usecase

import (
	"context"

	"github.com/JoseLora/fiberapp/internal/domain/entity"
)

type ProductFinderAll interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
}

type ProductFinderByID interface {
	FindByID(ctx context.Context, id string) (entity.Product, error)
}
