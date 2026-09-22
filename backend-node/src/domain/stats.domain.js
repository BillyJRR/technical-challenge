/**
 * Capa de Dominio: Validaciones y Entidades de Datos
 */

class StatsDomain {
    /**
     * Valida que la estructura del payload recibido de Go sea correcta.
     * @param {Array} matrices Array de objetos { name, data }
     */
    static validateMatricesPayload(matrices) {
        if (!Array.isArray(matrices) || matrices.length === 0) {
            throw new Error('El campo "matrices" debe ser un arreglo no vacío');
        }

        matrices.forEach((item, index) => {
            if (!item.data || !Array.isArray(item.data)) {
                throw new Error(`La matriz en el índice ${index} no contiene una propiedad "data" válida`);
            }

            const rows = item.data.length;
            if (rows === 0) {
                throw new Error(`La matriz "${item.name || index}" no puede estar vacía`);
            }

            const cols = item.data[0].length;
            item.data.forEach((row, rowIndex) => {
                if (!Array.isArray(row)) {
                    throw new Error(`La fila ${rowIndex} en la matriz "${item.name || index}" debe ser un arreglo`);
                }
                if (row.length !== cols) {
                    throw new Error(`Dimensiones inconsistentes en la matriz "${item.name || index}". Fila ${rowIndex} tiene longitud ${row.length}, se esperaba ${cols}`);
                }
                row.forEach((val) => {
                    if (typeof val !== 'number' || isNaN(val)) {
                        throw new Error(`Todos los elementos de la matriz deben ser números válidos. Valor inválido encontrado en "${item.name || index}"`);
                    }
                });
            });
        });

        return true;
    }
}

module.exports = StatsDomain;