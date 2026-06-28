# Moka

Full-stack project with Vite + Vue 3 + TypeScript frontend and Go + Gin backend.

## Tech Stack

### Frontend
- Vite
- Vue 3
- TypeScript
- Vue Router 4
- Pinia

### Backend
- Go
- Gin
- GORM
- PostgreSQL

## Getting Started

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Backend

```bash
cd backend
# Set environment variables (optional)
export DB_HOST=localhost
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=moka
export DB_PORT=5432

go run cmd/main.go
```

## API

- `GET /api/health` - Health check endpoint
