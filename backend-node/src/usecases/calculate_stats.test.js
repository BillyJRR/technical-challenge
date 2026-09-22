const test = require('node:test');
const assert = require('node:assert');
const { calculateStatsUseCase, isMatrixDiagonal } = require('./calculate_stats.usecase');

test('Cálculo correcto de Máximo, Mínimo, Suma y Promedio', () => {
    const payload = [
        {
            name: 'matriz_a',
            data: [
                [1, 2],
                [3, 4]
            ]
        }
    ];

    const stats = calculateStatsUseCase(payload);

    assert.strictEqual(stats.max, 4);
    assert.strictEqual(stats.min, 1);
    assert.strictEqual(stats.sum, 10);
    assert.strictEqual(stats.average, 2.5);
});

test('Detección correcta de Matriz Diagonal', () => {
    const diagonalMatrix = [
        [5, 0],
        [0, 2]
    ];
    const nonDiagonalMatrix = [
        [5, 1],
        [0, 2]
    ];

    assert.strictEqual(isMatrixDiagonal(diagonalMatrix), true);
    assert.strictEqual(isMatrixDiagonal(nonDiagonalMatrix), false);
});