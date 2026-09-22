const { Router } = require('express');
const StatsController = require('../controllers/stats.controller');
const verifyJWTMiddleware = require('../middlewares/jwt.middleware');

const router = Router();

router.post('/statistics', verifyJWTMiddleware, StatsController.getStatistics);

module.exports = router;