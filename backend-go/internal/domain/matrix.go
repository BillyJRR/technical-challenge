package domain

// Solicitud desde el Frontend
type MatrixRequest struct {
	Matrix [][]float64 `json:"matrix"`
}

// Estructura de matrices que se envía a Node.js
type NamedMatrix struct {
	Name string      `json:"name"`
	Data [][]float64 `json:"data"`
}

type NodePayload struct {
	Matrices []NamedMatrix `json:"matrices"`
}

// Respuesta de la API de Node.js
type DiagonalDetail struct {
	MatrixName string `json:"matrix_name"`
	IsDiagonal bool   `json:"is_diagonal"`
}

type NodeStatsData struct {
	Max               float64          `json:"max"`
	Min               float64          `json:"min"`
	Average           float64          `json:"average"`
	Sum               float64          `json:"sum"`
	HasDiagonalMatrix bool             `json:"has_diagonal_matrix"`
	DiagonalAnalysis  []DiagonalDetail `json:"diagonal_analysis"`
}

type NodeStatsResponse struct {
	Status string        `json:"status"`
	Data   NodeStatsData `json:"data"`
}

// Respuesta para el Frontend
type QRResult struct {
	Q [][]float64 `json:"q"`
	R [][]float64 `json:"r"`
}

type ProcessedMatrixResponse struct {
	OriginalMatrix [][]float64   `json:"original_matrix"`
	RotatedMatrix  [][]float64   `json:"rotated_matrix"`
	QR             QRResult      `json:"qr_factorization"`
	Statistics     NodeStatsData `json:"statistics"`
}

// Interfaz del Caso de Uso
type MatrixUseCase interface {
	Rotate90Clockwise(matrix [][]float64) [][]float64
	FactorizeQR(matrix [][]float64) ([][]float64, [][]float64, error)
	ProcessAndDelegate(matrix [][]float64, authToken string) (*ProcessedMatrixResponse, error)
}
