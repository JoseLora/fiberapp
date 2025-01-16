package handler

import (
	"github.com/JoseLora/fiberapp/internal/domain/entity"
	"github.com/JoseLora/fiberapp/internal/domain/usecase"
	"github.com/gofiber/fiber/v2"
)

type ProductAPI struct {
	findAll  usecase.ProductFinderAll
	findByID usecase.ProductFinderByID
}

func NewProductAPI(findAll usecase.ProductFinderAll,
	findByID usecase.ProductFinderByID) *ProductAPI {
	return &ProductAPI{
		findAll:  findAll,
		findByID: findByID,
	}
}

func (p *ProductAPI) FindAll(c *fiber.Ctx) ([]entity.Product, error) {
	return p.findAll.FindAll(c.Context())
}

func (p *ProductAPI) FindByID(c *fiber.Ctx, id string) (entity.Product, error) {
	return p.findByID.FindByID(c.Context(), id)
}
