package handler

import (
	"backend-go/internal/domain"

	"github.com/gofiber/fiber/v2"
)

type MatrixHandler struct {
	useCase domain.MatrixUseCase
}

func NewMatrixHandler(uc domain.MatrixUseCase) *MatrixHandler {
	return &MatrixHandler{useCase: uc}
}

func (h *MatrixHandler) ProcessMatrix(c *fiber.Ctx) error {
	var req domain.MatrixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Formato JSON inválido",
		})
	}

	authToken := c.Get("Authorization")

	response, err := h.useCase.ProcessAndDelegate(req.Matrix, authToken)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Matriz procesada correctamente",
		"data":    response,
	})
}
