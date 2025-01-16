package handler

import (
	"github.com/JoseLora/fiberapp/internal/domain/entity"
	"github.com/JoseLora/fiberapp/internal/domain/usecase"
	"github.com/gofiber/fiber/v2"
)

type ProductAPI interface {
	FindAll(c *fiber.Ctx) ([]entity.Product, error)
	FindByID(c *fiber.Ctx, id string) (entity.Product, error)
}

type ProductHandler struct {
	findAll  usecase.ProductFinderAll
	findByID usecase.ProductFinderByID
}

func NewProductAPI(findAll usecase.ProductFinderAll,
	findByID usecase.ProductFinderByID) *ProductHandler {
	return &ProductHandler{
		findAll:  findAll,
		findByID: findByID,
	}
}

func (p *ProductHandler) FindAll(c *fiber.Ctx) ([]entity.Product, error) {
	return p.findAll.FindAll(c.Context())
}

func (p *ProductHandler) FindByID(c *fiber.Ctx, id string) (entity.Product, error) {
	return p.findByID.FindByID(c.Context(), id)
}
