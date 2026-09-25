# 14 — Deep Analysis Findings

## Scope

The uploaded ZIP was inspected at source level: Go application code, routes, repositories, PostgreSQL schema, React source, QA tree, Jenkinsfile, Makefile, Git metadata, YARA rules and security configuration.

## Major packages

The project contains dedicated modules for attachment analysis, BEC, campaign correlation, decision, detonation, DNS reputation, domain analysis, email authentication, headers, hashing, IOC, macro analysis, MITRE, PDF, PhishTank, QR, risk, sandbox, sender, Sigma, SMTP, Splunk, STIX, thread hijack, threat feeds, threat intelligence, timeline, URL analysis, VirusTotal, WebSocket, WHOIS, YARA and ZIP analysis.

## Strong architecture point

The central email pipeline collects many evidence sources before risk calculation, creating an explainable evidence → score → decision workflow.

## WebSocket verified

It is actually implemented: `routes/websocket_routes.go`, `internal/websocket/websocket_server.go` and the email pipeline use Gorilla WebSocket and a broadcast channel.

## SMTP verified

The backend starts an SMTP server on port 2525 and sends parsed messages into the processing pipeline.

## Sandbox finding

There are static/content checks, rule-based behavior checks and a Docker execution path. Some detonation results are inferred from file type/content, so they are not equivalent to real runtime telemetry.

## Email-authentication finding

SPF/DKIM/DMARC logic is simulated by sender-domain logic. It should not be presented as full standards-compliant validation.

## Threat-intelligence finding

The platform is hybrid: VirusTotal integration + local feed + local reputation + hard-coded demo intelligence.

## Authentication finding

JWT is present, but the source has a hard-coded signing secret, direct password handling and inconsistent protected-route coverage. Production hardening is required.

## Frontend finding

The frontend is React + Vite + React Router, not Next.js.

## CI finding

Jenkins and Makefile automation are present. The active Jenkins stages build backend/frontend and run Go tests. The Playwright stage is present but commented out.

## QA finding

The QA tree is substantial and includes API, integration, E2E, smoke, load, performance and UI assets. Existing reports are historical artifacts; they are not proof of current passing status.

## Git history finding

Git history contains commits mentioning ransomware and deployment work, but the current tree does not contain a ransomware-named package. Current documentation is based on the current source tree.

## Generated artifacts

The archive includes node_modules, Playwright artifacts, compiled frontend output and test files. These are not core application source.

## Final characterization

> A full-stack defensive SOC investigation platform combining deterministic phishing/file detection, threat intelligence, IOC correlation, risk scoring, MITRE mapping, case management, sandbox-style analysis and real-time analyst alerting.

It should not be described as a complete commercial sandbox, a fully autonomous malware-analysis system, an ML phishing classifier or a fully hardened enterprise IAM platform.
