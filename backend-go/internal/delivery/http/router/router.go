package router

import (
	"backend-go/internal/config"
	"backend-go/internal/delivery/http/handler"
	"backend-go/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, matrixHandler *handler.MatrixHandler, authHandler *handler.AuthHandler) {
	// Grupo API v1
	api := app.Group("/api/v1")

	// Ruta de Health Check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "healthy",
			"service": "Go Fiber Matrix API",
		})
	})

	api.Post("/auth/token", authHandler.GetToken)
	api.Post("/process-matrix", middleware.JWTMiddleware(cfg), matrixHandler.ProcessMatrix)
}
