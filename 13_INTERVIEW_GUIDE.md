# 13 — Interview Guide

## 30-second answer

> I built a full-stack SOC platform for phishing and suspicious-file investigation. The backend is Go and the UI is React/Vite. It combines parsing, YARA-style detection, domain analysis, threat intelligence, hashing, sandbox analysis, deterministic risk scoring and MITRE mapping, then stores results in PostgreSQL and provides dashboards, cases and real-time alerts.

## Email pipeline

```text
Email → Parse → Extract → Detect → Threat Intel → Risk → Decision → Store → Alert
```

## Why Go?

Simple concurrency, fast compilation, strong networking support and suitability for security tooling.

## Why PostgreSQL?

The platform contains relational users, emails, cases, alerts, IOCs and sandbox data.

## Is it AI?

The current core is mainly deterministic/rule-based. It is not correct to call it an ML phishing classifier. ML could be added later.

## Explain sandbox

The project has a Docker-based analysis path with resource limits and a timeout, plus rule/content behavior analysis. It is not a complete commercial malware sandbox.

## Explain WebSocket

After analysis, selected result fields are JSON-encoded and sent through an in-memory broadcast channel to connected browser clients.

## Explain IOC

An IOC is an indicator such as a URL, domain, IP, hash or email address that can support compromise investigation.

## Explain MITRE

The platform maps selected findings to ATT&CK-style techniques so analysts can understand attack behavior in a common framework.

## What would you improve?

- password hashing
- secret management
- complete RBAC
- route authorization audit
- upload hardening
- WebSocket origin restrictions
- stronger sandbox isolation
- real SPF/DKIM/DMARC validation
- richer threat-intelligence lifecycle
- risk calibration
- database migrations/constraints
- security-focused CI tests

## Strong closing statement

> The key design idea is to combine many explainable security signals into one investigation workflow rather than relying on a single detector.
