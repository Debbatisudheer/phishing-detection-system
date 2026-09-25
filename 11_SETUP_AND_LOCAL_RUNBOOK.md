# 11 — Setup and Local Runbook

## Prerequisites

- Go 1.24.x-compatible toolchain
- Node.js + npm
- PostgreSQL
- Docker Desktop
- Git

## Database

Create `phishing_platform` and apply `database/schema.sql`.

Configure either `DATABASE_URL` or DB_HOST/DB_PORT/DB_USER/DB_PASSWORD/DB_NAME.

## Backend

From project root:

```powershell
go mod download
go build ./...
go test ./... -v
go vet ./...
go run .
```

Default backend: `http://localhost:8081`

Health: `http://localhost:8081/health`

## Frontend

```powershell
cd frontend
npm install
npm run dev
```

Set `VITE_API_URL` if the backend is not at the default URL.

## SMTP

The backend starts an SMTP server on port `2525`.

## Recommended startup order

```text
PostgreSQL → Docker → Go backend → React frontend → optional SMTP test client
```

## Docker sandbox

The source expects an image named `phishing-sandbox` for its Docker execution path.

## Troubleshooting

DB failure: verify PostgreSQL, database, credentials and environment variables.

Frontend API failure: verify backend, `VITE_API_URL` and CORS.

WebSocket failure: verify backend `/ws` and browser/network logs.

Sandbox failure: verify Docker, expected image, mounted paths and timeout.

Never run unknown malware on a normal workstation.
