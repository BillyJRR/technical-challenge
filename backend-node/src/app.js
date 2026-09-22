const express = require('express');
const cors = require('cors');
const statsRoutes = require('./routes/stats.routes');

const app = express();

app.use(cors());
app.use(express.json());

// Rutas de la API
app.use('/api/v1', statsRoutes);

module.exports = app;