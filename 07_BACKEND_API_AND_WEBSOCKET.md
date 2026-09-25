# 07 — Backend API and WebSocket

## Startup

`main.go` loads `.env`, connects PostgreSQL, starts background workers, starts SMTP, starts the WebSocket handler, registers routes, configures CORS and starts HTTP on `PORT` or default `8081`.

## Key endpoints

Authentication:
- `POST /api/register`
- `POST /api/login`

Analysis:
- `POST /api/analyze-email`
- `POST /api/analyze-file`

Other route groups provide findings, search, cases, incidents, campaigns, threat intelligence, sandbox reports, IOC graphs, MITRE statistics, reports and system health.

Health:
- `GET /health`

WebSocket:
- `GET /ws`

## JWT

JWT claims include username, role and expiration. The frontend stores the token and Axios sends it as a Bearer token.

## WebSocket flow

```text
analysis
 ↓
JSON event
 ↓
Broadcast channel
 ↓
WebSocket worker
 ↓
connected clients
```

## Current security gaps

Source review found a hard-coded JWT secret, direct password handling, inconsistent JWT middleware coverage and a WebSocket upgrader that accepts all origins. These are production hardening items.
