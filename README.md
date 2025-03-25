# Stock Tracker

Una aplicación web moderna para el seguimiento y análisis de acciones bursátiles, construida con una arquitectura de microservicios.

## Características

- Seguimiento en tiempo real de acciones bursátiles
- Recomendaciones de mejores acciones
- Interfaz de usuario moderna y responsive
- API RESTful segura
- Manejo de CORS y seguridad implementada

## Tecnologías Utilizadas

### Backend
- Go (Golang)
- Gin Web Framework
- PostgreSQL (Base de datos)
- Docker

### Frontend
- Vue.js 3
- TypeScript
- Tailwind CSS
- Pinia (Estado)
- Vue Router

## Prerrequisitos

- Go 1.21 o superior
- Node.js 18 o superior
- Docker y Docker Compose
- PostgreSQL

## Instalación

### Backend

1. Navega al directorio del backend:
```bash
cd Backend
```

2. Copia el archivo de variables de entorno:
```bash
cp .env.example .env
```

3. Instala las dependencias:
```bash
go mod download
```

4. Inicia el servidor:
```bash
go run main.go
```

El servidor estará disponible en `http://localhost:9090`

### Frontend

1. Navega al directorio del frontend:
```bash
cd Frontend/stock-tracker
```

2. Instala las dependencias:
```bash
npm install
```

3. Inicia el servidor de desarrollo:
```bash
npm run dev
```

La aplicación estará disponible en `http://localhost:5173`

## Usando Docker

Para ejecutar todo el stack con Docker:

```bash
docker-compose up -d
```

## API Endpoints

### Stocks
- `GET /stocks` - Obtiene la lista de acciones
- `GET /stocks/recommendations` - Obtiene recomendaciones de mejores acciones

## Configuración de Seguridad

El proyecto incluye:
- Middleware de seguridad
- Configuración CORS
- Variables de entorno para credenciales
- Manejo seguro de tokens

## Scripts Disponibles

### Backend
- `go run main.go` - Inicia el servidor de desarrollo
- `go test ./...` - Ejecuta las pruebas

### Frontend
- `npm run dev` - Inicia el servidor de desarrollo
- `npm run build` - Construye para producción
- `npm run lint` - Ejecuta el linter
- `npm run format` - Formatea el código

###Creador
Samir Esteban Gonzalez 