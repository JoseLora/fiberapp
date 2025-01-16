package usecase

import (
	"context"

	"github.com/JoseLora/fiberapp/internal/domain/entity"
	"github.com/JoseLora/fiberapp/internal/domain/repository"
	"github.com/JoseLora/fiberapp/internal/domain/usecase"
)

// ProductFinderAll is an implementation of the ProductFinderAll use case.
// It uses a repository to find all products.
type ProductFinderAll struct {
	repository repository.Product
}

// NewProductFinderAll creates a new instance of ProductFinderAllImpl.
// It takes a repository as a parameter and returns a ProductFinderAll use case.
func NewProductFinderAll(repository repository.Product) usecase.ProductFinderAll {
	return &ProductFinderAll{
		repository: repository,
	}
}

// FindAll retrieves all products from the repository.
// It takes a context as a parameter and returns a slice of products and an error.
func (p *ProductFinderAll) FindAll(ctx context.Context) ([]entity.Product, error) {
	return p.repository.FindAll(ctx)
}

// ProductFinderByID is an implementation of the ProductFinderByID use case.
// It uses a repository to find a product by its ID.
type ProductFinderByID struct {
	repository repository.Product
}

// NewProductFinderByID creates a new instance of ProductFinderByIDImpl.
// It takes a repository as a parameter and returns a ProductFinderByID use case.
func NewProductFinderByID(repository repository.Product) usecase.ProductFinderByID {
	return &ProductFinderByID{
		repository: repository,
	}
}

// FindByID retrieves a product by its ID from the repository.
// It takes a context and an ID as parameters and returns a product and an error.
func (p *ProductFinderByID) FindByID(ctx context.Context, id string) (entity.Product, error) {
	return p.repository.FindByID(ctx, id)
}
