package repository

import (
	"context"

	"github.com/JoseLora/fiberapp/internal/domain/entity"
)

type Product interface {
	FindAll(ctx context.Context) ([]entity.Product, error)
	FindByID(ctx context.Context, id string) (entity.Product, error)
}
