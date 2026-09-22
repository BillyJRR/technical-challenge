# Desafío Técnico - Procesamiento de Matrices y Estadísticas

Sistema distribuido full-stack para el procesamiento matemático de matrices y cálculo estadístico. Construido bajo **Clean Architecture**, comunicación entre servicios restringida con **JWT** y orquestación local vía **Docker Compose**.

---

## Arquitectura del Sistema

El proyecto está dividido en tres capas/servicios:

1. **Frontend (Nginx / S3):** Interfaz web interactiva para la creación y visualización de resultados.
2. **Backend Principal (Go + Fiber):** 
   - Genera tokens JWT de sesión.
   - Aplica rotación de matriz ($90^\circ, 180^\circ, 270^\circ$).
   - Realiza la descomposición $QR$ mediante el algoritmo de Gram-Schmidt.
   - Orquesta la llamada HTTP hacia la API de Node.js.
3. **Backend Estadístico (Node.js + Express):**
   - Válida peticiones mediante middleware JWT.
   - Calcula mínimo, máximo, promedio, suma total y verifica si la matriz es diagonal.

---

## Ejecución en Local (Docker Compose)

### Requisitos previos
- Docker Desktop instalado y en ejecución.

### Pasos
1. Clonar el repositorio:
   git clone [https://github.com/tu-usuario/tu-repositorio.git]
   cd tu-repositorio
2. Clonar el repositorio:
    docker-compose up --build
3. Abrir la aplicación en el navegador:
    Frontend: http://localhost
    API Go: http://localhost:8080
    API Node.js: http://localhost:3000