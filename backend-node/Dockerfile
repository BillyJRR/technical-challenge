FROM node:20-alpine

WORKDIR /usr/src/app

# Copiar archivos de definición de dependencias
COPY package*.json ./

# Instalar dependencias fijas para producción
RUN npm ci --only=production

# Copiar el resto del código
COPY . .

EXPOSE 3000

USER node

CMD ["npm", "start"]