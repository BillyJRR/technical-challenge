package usecase

import (
	"testing"
)

func TestRotate90Clockwise(t *testing.T) {
	uc := &matrixUseCase{}
	input := [][]float64{
		{1, 2},
		{3, 4},
	}
	expected := [][]float64{
		{3, 1},
		{4, 2},
	}

	result := uc.Rotate90Clockwise(input)

	for i := range expected {
		for j := range expected[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("Se esperaba %f en [%d][%d], se obtuvo %f", expected[i][j], i, j, result[i][j])
			}
		}
	}
}
