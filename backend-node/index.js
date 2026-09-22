const app = require('./src/app');
const config = require('./src/config/env.config');

app.listen(config.PORT, () => {
    console.log(`Servidor Node.js Express escuchando en el puerto ${config.PORT}`);
});