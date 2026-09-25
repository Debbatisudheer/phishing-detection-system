# 01 — Architecture and Data Flow

## Layers

```text
React/Vite UI
   ↓
Go HTTP API
   ↓
Security analysis modules
   ↓
PostgreSQL repositories
   ↓
WebSocket + reports + IOC/STIX/Sigma/Splunk outputs
```

## Backend structure

- `routes/`: endpoint registration
- `internal/api/`: feature API handlers
- `internal/`: security/detection logic
- `database/`: PostgreSQL connection/repositories
- `models/`: shared models

## Email path

```text
SMTP 2525 or POST /api/analyze-email
  ↓
parser
  ↓
pipeline.ProcessEmail
  ↓
YARA/header/BEC/thread/auth/sender/domain/TI/attachments/etc.
  ↓
risk.CalculateRisk
  ↓
decision.MakeDecision
  ↓
DB + reports + WebSocket
```

## File path

```text
POST /api/analyze-file
  ↓
uploads/
  ↓
file-specific analyzer
  ↓
YARA/hash/TI/PDF/macro/ZIP/QR
  ↓
risk + MITRE + verdict
  ↓
DB
  ↓
sandbox job when required
```

## WebSocket

`GET /ws` is actually implemented using Gorilla WebSocket. Connected clients are kept in memory. Analysis results are JSON-encoded and sent through a broadcast channel.

## CI

Jenkins stages: checkout, environment check, database verification, schema import, backend build, frontend build and backend tests. The Playwright stage exists but is commented out.
