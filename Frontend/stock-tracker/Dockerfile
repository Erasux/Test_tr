FROM node:20-alpine

WORKDIR /app

# Copiar archivos de dependencias
COPY package*.json ./

# Instalar dependencias
RUN npm install

# Copiar el resto del código
COPY . .

# Construir la aplicación
RUN npm run build

# Instalar servidor ligero para producción
RUN npm install -g serve

# Exponer el puerto
EXPOSE 5173

# Comando para servir la aplicación
CMD ["npm", "run", "start"] 