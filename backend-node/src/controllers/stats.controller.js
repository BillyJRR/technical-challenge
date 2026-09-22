const { calculateStatsUseCase } = require('../usecases/calculate_stats.usecase');
const StatsDomain = require('../domain/stats.domain');

class StatsController {
    static getStatistics(req, res) {
        try {
            const { matrices } = req.body;

            // Validación de dominio
            StatsDomain.validateMatricesPayload(matrices);

            // Ejecución del caso de uso
            const result = calculateStatsUseCase(matrices);

            return res.status(200).json({
                status: 'success',
                data: result
            });
        } catch (error) {
            return res.status(400).json({
                status: 'error',
                message: error.message
            });
        }
    }
}

module.exports = StatsController;