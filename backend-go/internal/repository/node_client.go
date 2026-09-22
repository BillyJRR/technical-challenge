package repository

import (
	"backend-go/internal/domain"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type nodeClient struct {
	nodeURL   string
	jwtSecret string
}

func NewNodeClient(nodeURL string, jwtSecret string) *nodeClient {
	return &nodeClient{nodeURL: nodeURL, jwtSecret: jwtSecret}
}

func (c *nodeClient) GetStatistics(payload domain.NodePayload, authToken string) (*domain.NodeStatsResponse, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error al serializar payload para node: %w", err)
	}

	agent := fiber.Post(c.nodeURL)
	agent.Body(body)
	agent.ContentType("application/json")
	agent.Set("Authorization", authToken)
	agent.Timeout(5 * time.Second)

	statusCode, responseBody, errs := agent.Bytes()
	if len(errs) > 0 {
		return nil, fmt.Errorf("error de red al conectar con Node.js API: %v", errs[0])
	}

	if statusCode != fiber.StatusOK {
		return nil, fmt.Errorf("Node.js API devolvió status: %d", statusCode)
	}

	var statsRes domain.NodeStatsResponse
	if err := json.Unmarshal(responseBody, &statsRes); err != nil {
		return nil, fmt.Errorf("error al deserializar respuesta de Node.js: %w", err)
	}

	return &statsRes, nil
}
