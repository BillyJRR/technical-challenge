package domain

import "errors"

var (
	ErrEmptyMatrix       = errors.New("la matriz no puede estar vacía")
	ErrInvalidDimensions = errors.New("las filas de la matriz deben tener la misma dimensión")
	ErrUnauthorized      = errors.New("token de autorización no válido o ausente")
)
