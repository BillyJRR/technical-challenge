const jwt = require('jsonwebtoken');
const config = require('../config/env.config');

function verifyJWTMiddleware(req, res, next) {
    const authHeader = req.headers['authorization'];

    if (!authHeader) {
        return res.status(401).json({
            status: 'error',
            message: 'Acceso denegado: No se proporcionó el token de autorización'
        });
    }

    const token = authHeader.split(' ')[1]; // Formato "Bearer <token>"
    if (!token) {
        return res.status(401).json({
            status: 'error',
            message: 'Formato de token inválido'
        });
    }

    try {
        const verified = jwt.verify(token, config.JWT_SECRET);
        req.user = verified;
        next(); // Permite el paso al controlador si es válido
    } catch (error) {
        return res.status(403).json({
            status: 'error',
            message: 'Token inválido o expirado'
        });
    }
}

module.exports = verifyJWTMiddleware;