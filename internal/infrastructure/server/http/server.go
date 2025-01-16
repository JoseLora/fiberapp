package http

import (
	"github.com/JoseLora/fiberapp/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	engine *fiber.App
}

func NewServer(productHandler *handler.ProductHandler) *Server {
	app := fiber.New()
	app.Get("/products", func(c *fiber.Ctx) error {
		if products, error := productHandler.FindAll(c); error != nil {
			return error
		} else {
			return c.JSON(products)
		}
	})
	app.Get("/products/:id", func(c *fiber.Ctx) error {
		if product, error := productHandler.FindByID(c, c.Params("id")); error != nil {
			return error
		} else {
			return c.JSON(product)
		}
	})

	return &Server{
		engine: app,
	}
}

func (s *Server) Start() error {
	return s.engine.Listen(":8080")
}

func (s *Server) Stop() error {
	return s.engine.Shutdown()
}
