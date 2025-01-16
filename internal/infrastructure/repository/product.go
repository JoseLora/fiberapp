package repository

import (
	"context"
	"errors"
	"sync"

	"github.com/JoseLora/fiberapp/internal/domain/entity"
	"github.com/JoseLora/fiberapp/internal/domain/repository"
)

var ErrProductNotFound = errors.New("product not found")

type ProductInMemory struct {
	DB sync.Map
}

type ProductInMemoryEntity struct {
	ID          string
	Name        string
	Description string
	Discount    float64
}

func NewProductInMemory() repository.Product {
	repo := &ProductInMemory{}

	repo.DB.Store("1", ProductInMemoryEntity{ID: "1", Name: "Product 1", Description: "Description 1", Discount: 0.1})
	repo.DB.Store("2", ProductInMemoryEntity{ID: "2", Name: "Product 2", Description: "Description 2", Discount: 0.2})
	repo.DB.Store("3", ProductInMemoryEntity{ID: "3", Name: "Product 3", Description: "Description 3", Discount: 0.3})

	return repo
}

func (repo *ProductInMemory) FindAll(ctx context.Context) ([]entity.Product, error) {
	products := make([]entity.Product, 0)

	repo.DB.Range(func(key, value interface{}) bool {
		product := value.(ProductInMemoryEntity)
		// TODO use a mapper
		products = append(products, entity.Product{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Discount:    product.Discount,
		})
		return true
	})

	return products, nil
}

func (repo *ProductInMemory) FindByID(ctx context.Context, id string) (entity.Product, error) {
	if product, ok := repo.DB.Load(id); ok {
		p := product.(ProductInMemoryEntity)
		// TODO use a mapper
		pEntity := entity.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Discount:    p.Discount,
		}
		return pEntity, nil
	} else {
		return entity.Product{}, ErrProductNotFound
	}
}
