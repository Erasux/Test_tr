FROM golang:1.21-alpine

WORKDIR /app

# Instalar dependencias del sistema
RUN apt-get update && apt-get install -y git

# Copiar archivos de dependencias
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto
EXPOSE 9090

# Comando para ejecutar la aplicación
CMD ["./main"] 