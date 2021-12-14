package router

import (
	"github.com/OrlandoRomo/go-ambassador/pkg/insfrastructure/middleware"
	"github.com/OrlandoRomo/go-ambassador/pkg/interface/controller"
	"github.com/gofiber/fiber/v2"
)

func SetProductRoutes(r *fiber.Router, c *controller.AppController) {
	product := *r
	product.Get("/products/", middleware.AuthMiddleware, c.Product.GetProducts)
}