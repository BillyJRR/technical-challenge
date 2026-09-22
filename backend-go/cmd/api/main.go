package main

import (
	"backend-go/internal/config"
	"backend-go/internal/delivery/http/handler"
	"backend-go/internal/delivery/http/router"
	"backend-go/internal/repository"
	"backend-go/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.LoadConfig()

	// Inyección de dependencias
	nodeClient := repository.NewNodeClient(cfg.NodeServiceURL, cfg.JWTSecret)
	matrixUseCase := usecase.NewMatrixUseCase(nodeClient)

	matrixHandler := handler.NewMatrixHandler(matrixUseCase)
	authHandler := handler.NewAuthHandler(cfg)

	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Registrar Rutas
	router.SetupRoutes(app, cfg, matrixHandler, authHandler)

	log.Printf("Servidor Go Fiber escuchando en el puerto %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
