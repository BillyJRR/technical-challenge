/**
 * Verifica si una matriz de números es diagonal.
 * Condición: Debe ser cuadrada (N x N) y todos los elementos a_ij con i != j deben ser 0.
 */
function isMatrixDiagonal(matrix) {
    if (!Array.isArray(matrix) || matrix.length === 0) return false;

    const rows = matrix.length;
    const cols = matrix[0].length;

    // Una matriz diagonal debe ser cuadrada
    if (rows !== cols) return false;

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            if (i !== j && Math.abs(matrix[i][j]) > 1e-6) { // Tolerancia para flotantes
                return false;
            }
        }
    }
    return true;
}

/**
 * Calcula las estadísticas globales sobre un conjunto de matrices.
 * @param {Array<{name: string, data: Array<Array<number>>}>} matrices 
 */
function calculateStatsUseCase(matrices) {
    if (!Array.isArray(matrices) || matrices.length === 0) {
        throw new Error('Debe proporcionar al menos una matriz');
    }

    let max = -Infinity;
    let min = Infinity;
    let sum = 0;
    let count = 0;
    let hasDiagonalMatrix = false;
    const diagonalAnalysis = [];

    matrices.forEach((item) => {
        const { name, data } = item;

        // Evaluar si esta matriz específica es diagonal
        const isDiag = isMatrixDiagonal(data);
        if (isDiag) {
            hasDiagonalMatrix = true;
        }

        diagonalAnalysis.push({
            matrix_name: name || 'unnamed_matrix',
            is_diagonal: isDiag
        });

        // Aplanar y recorrer elementos para las métricas estadísticas
        data.forEach(row => {
            row.forEach(val => {
                if (val > max) max = val;
                if (val < min) min = val;
                sum += val;
                count++;
            });
        });
    });

    if (count === 0) {
        throw new Error('Las matrices proporcionadas están vacías');
    }

    const average = sum / count;

    return {
        max: Number(max.toFixed(4)),
        min: Number(min.toFixed(4)),
        average: Number(average.toFixed(4)),
        sum: Number(sum.toFixed(4)),
        has_diagonal_matrix: hasDiagonalMatrix,
        diagonal_analysis: diagonalAnalysis
    };
}

module.exports = {
    calculateStatsUseCase,
    isMatrixDiagonal
};