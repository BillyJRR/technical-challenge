// URL de la API de Go desplegada en Render (o localhost para desarrollo)
const API_GO_URL = window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1'
    ? 'http://localhost:8080/api/v1/process-matrix'
    : 'https://tu-api-go-en-render.onrender.com/api/v1/process-matrix';

// Inicializar la matriz por defecto al cargar la página
document.addEventListener('DOMContentLoaded', generateMatrixInput);

function generateMatrixInput() {
    const rows = parseInt(document.getElementById('rows').value) || 3;
    const cols = parseInt(document.getElementById('cols').value) || 3;
    const container = document.getElementById('matrix-input-container');

    container.style.gridTemplateColumns = `repeat(${cols}, 60px)`;
    container.innerHTML = '';

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            const input = document.createElement('input');
            input.type = 'number';
            input.className = 'cell-input';
            input.value = Math.floor(Math.random() * 10) + 1; // Valor por defecto aleatorio
            input.dataset.row = i;
            input.dataset.col = j;
            container.appendChild(input);
        }
    }

    document.getElementById('process-btn').style.display = 'inline-block';
}

function getMatrixData() {
    const rows = parseInt(document.getElementById('rows').value);
    const cols = parseInt(document.getElementById('cols').value);
    const matrix = Array.from({ length: rows }, () => Array(cols).fill(0));
    const inputs = document.querySelectorAll('.cell-input');

    inputs.forEach(input => {
        const r = parseInt(input.dataset.row);
        const c = parseInt(input.dataset.col);
        matrix[r][c] = parseFloat(input.value) || 0;
    });

    return matrix;
}

let jwtToken = '';

// Obtener token JWT al cargar la aplicación
async function authenticateJWT() {
    const AUTH_URL = API_GO_URL.replace('/process-matrix', '/auth/token');
    try {
        const res = await fetch(AUTH_URL, { method: 'POST' });
        if (!res.ok) throw new Error('No se pudo autenticar con el servidor');

        const data = await res.json();
        if (data.token) {
            jwtToken = data.token;
            console.log('JWT autenticado con éxito');
        }
    } catch (e) {
        console.error('Error al obtener token JWT:', e);
    }
}

document.addEventListener('DOMContentLoaded', () => {
    generateMatrixInput();
    authenticateJWT();
});

async function processMatrix() {
    const matrixData = getMatrixData();
    const loading = document.getElementById('loading');
    const resultsContainer = document.getElementById('results-container');

    loading.style.display = 'block';
    resultsContainer.style.display = 'none';

    try {
        const response = await fetch(API_GO_URL, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${jwtToken}` },
            body: JSON.stringify({ matrix: matrixData })
        });

        if (!response.ok) {
            throw new Error(`Error en el servidor: ${response.statusText}`);
        }

        const res = await response.json();
        displayResults(res.data);
    } catch (err) {
        alert('Ocurrió un error al procesar la matriz: ' + err.message);
    } finally {
        loading.style.display = 'none';
    }
}

function displayResults(data) {
    document.getElementById('results-container').style.display = 'block';

    // Renderizar matrices de la API de Go
    renderMatrixTable('rotated-matrix-display', data.rotated_matrix);
    renderMatrixTable('q-matrix-display', data.qr_factorization.q);
    renderMatrixTable('r-matrix-display', data.qr_factorization.r);

    // Renderizar estadísticas de la API de Node.js
    const stats = data.statistics;
    document.getElementById('stat-max').innerText = stats.max;
    document.getElementById('stat-min').innerText = stats.min;
    document.getElementById('stat-avg').innerText = stats.average;
    document.getElementById('stat-sum').innerText = stats.sum;
    document.getElementById('stat-diag').innerText = stats.has_diagonal_matrix ? 'SÍ' : 'NO';
}

function renderMatrixTable(containerId, matrix) {
    const container = document.getElementById(containerId);
    if (!matrix || matrix.length === 0) {
        container.innerHTML = '<p>Sin datos</p>';
        return;
    }

    let html = '<table class="matrix-table">';
    matrix.forEach(row => {
        html += '<tr>';
        row.forEach(val => {
            html += `<td>${typeof val === 'number' ? val.toFixed(2) : val}</td>`;
        });
        html += '</tr>';
    });
    html += '</table>';
    container.innerHTML = html;
}