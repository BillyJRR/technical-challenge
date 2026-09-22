package usecase

import (
	"backend-go/internal/domain"
	"math"
)

type NodeClient interface {
	GetStatistics(payload domain.NodePayload, authToken string) (*domain.NodeStatsResponse, error)
}

type matrixUseCase struct {
	nodeClient NodeClient
}

func NewMatrixUseCase(nc NodeClient) domain.MatrixUseCase {
	return &matrixUseCase{nodeClient: nc}
}

// Rotación 90° en sentido horario: Transponer y luego invertir cada fila
func (u *matrixUseCase) Rotate90Clockwise(matrix [][]float64) [][]float64 {
	rows := len(matrix)
	if rows == 0 {
		return [][]float64{}
	}
	cols := len(matrix[0])

	// Crear nueva matriz transpusta de dimensión cols x rows
	rotated := make([][]float64, cols)
	for i := range rotated {
		rotated[i] = make([]float64, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = matrix[i][j]
		}
	}
	return rotated
}

// Factorización QR mediante Gram-Schmidt Modificado
func (u *matrixUseCase) FactorizeQR(matrix [][]float64) ([][]float64, [][]float64, error) {
	m := len(matrix)
	if m == 0 {
		return nil, nil, domain.ErrEmptyMatrix
	}
	n := len(matrix[0])

	// Inicializar Q (m x n) y R (n x n)
	Q := make([][]float64, m)
	for i := range Q {
		Q[i] = make([]float64, n)
	}
	R := make([][]float64, n)
	for i := range R {
		R[i] = make([]float64, n)
	}

	// Copia de columnas de A para trabajar
	V := make([][]float64, m)
	for i := 0; i < m; i++ {
		V[i] = make([]float64, n)
		copy(V[i], matrix[i])
	}

	for j := 0; j < n; j++ {
		// Calcular norma de la columna j
		var norm float64
		for i := 0; i < m; i++ {
			norm += V[i][j] * V[i][j]
		}
		norm = math.Sqrt(norm)

		R[j][j] = norm

		if norm > 1e-10 {
			for i := 0; i < m; i++ {
				Q[i][j] = V[i][j] / norm
			}
		}

		for k := j + 1; k < n; k++ {
			// Producto escalar Q[:, j] . V[:, k]
			var dot float64
			for i := 0; i < m; i++ {
				dot += Q[i][j] * V[i][k]
			}
			R[j][k] = dot
			for i := 0; i < m; i++ {
				V[i][k] -= dot * Q[i][j]
			}
		}
	}

	return Q, R, nil
}

func (u *matrixUseCase) ProcessAndDelegate(matrix [][]float64, authToken string) (*domain.ProcessedMatrixResponse, error) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, domain.ErrEmptyMatrix
	}

	rotated := u.Rotate90Clockwise(matrix)
	qMatrix, rMatrix, err := u.FactorizeQR(matrix)
	if err != nil {
		return nil, err
	}

	// Construir Payload para enviar a Node.js
	nodePayload := domain.NodePayload{
		Matrices: []domain.NamedMatrix{
			{Name: "rotated_matrix", Data: rotated},
			{Name: "q_matrix", Data: qMatrix},
			{Name: "r_matrix", Data: rMatrix},
		},
	}

	statsResponse, err := u.nodeClient.GetStatistics(nodePayload, authToken)
	if err != nil {
		return nil, err
	}

	return &domain.ProcessedMatrixResponse{
		OriginalMatrix: matrix,
		RotatedMatrix:  rotated,
		QR: domain.QRResult{
			Q: qMatrix,
			R: rMatrix,
		},
		Statistics: statsResponse.Data,
	}, nil
}
