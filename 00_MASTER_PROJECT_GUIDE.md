# 00 — Master Project Guide

## What is this project?

A full-stack **Phishing Detection and SOC Investigation Platform**. It receives suspicious emails/files, runs security checks, calculates risk, chooses a verdict, stores results and presents them to analysts.

## Technology

- Backend: Go
- Frontend: React + Vite
- Database: PostgreSQL
- Authentication: JWT
- Real-time: Gorilla WebSocket
- Email ingestion: SMTP on port 2525
- Detection: YARA and rule-based analyzers
- Threat intelligence: VirusTotal, local feeds and reputation data
- Sandbox path: Docker
- CI: Jenkins

## Main flow

```text
Email/File
  ↓
Parse / extract
  ↓
Detection engines
  ↓
Threat intelligence
  ↓
Collect findings
  ↓
Risk score
  ↓
Risk level + verdict
  ↓
PostgreSQL
  ↓
Dashboard / WebSocket / reports
```

## Email analysis

The central pipeline combines URL extraction, header checks, display-name spoofing, BEC detection, thread-hijack detection, email-authentication checks, sender checks, campaign checks, QR checks, domain analysis, threat intelligence, attachments, YARA, PDF/macro/ZIP processing, hashing, VirusTotal, risk scoring, MITRE mapping and reporting.

## File analysis

The source has paths for DOCM/XLSM, PS1, ZIP, PDF and image/QR analysis. SHA-256 and VirusTotal checks are also present. Suspicious file types can create sandbox jobs.

## SOC features

Dashboards, cases, case notes, incidents, campaigns, investigation, search, threat hunting, alerts, IOC graphs/trends, MITRE dashboards and sandbox reports are present.

## Database

The supplied schema has 15 tables: `alerts`, `analysis_results`, `analyst_notes`, `campaigns`, `case_notes`, `cases`, `docker_reports`, `emails`, `ioc_correlation`, `mitre_events`, `sandbox_jobs`, `sandbox_reports`, `threat_intel_cache`, `url_reputation`, `users`.

## Important accuracy

Some components are simulated or rule-based. The current SPF/DKIM/DMARC component is simulation-based, and some sandbox/detonation behavior is inferred from content or filenames. Do not describe these as equivalent to full commercial security products.

## Repository state

The archive is on `qa-automation` at commit `10359bed0bb2c1a27878c9a3bec73cc5e6047cbb`.

## Interview answer

> I built a Go and React SOC platform that analyzes phishing emails and suspicious files using multiple detection engines, threat intelligence, risk scoring and MITRE mapping, then stores and visualizes investigation results for analysts.
