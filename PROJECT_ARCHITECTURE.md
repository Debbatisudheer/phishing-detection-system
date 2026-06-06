# Phishing Detection & SOC Investigation Platform

## 1. Project Overview

A full-stack Security Operations Center (SOC) platform designed to detect, analyze, investigate, and respond to phishing emails, malicious files, and threat intelligence indicators.

Technology Stack:

* Backend: Go (Golang)
* Frontend: React
* Database: PostgreSQL
* Authentication: JWT
* Real-Time Alerts: WebSocket
* Threat Intelligence: VirusTotal

---

# 2. Core Modules

## Authentication

Features:

* User Registration
* User Login
* JWT Token Generation
* JWT Validation Middleware
* Protected APIs

---

## Email Security Engine

Features:

* Email Parsing
* Email Header Analysis
* SPF Validation
* DKIM Validation
* DMARC Validation
* Sender Reputation Analysis
* Domain Age Analysis
* URL Extraction
* URL Reputation Checks
* Phishing Detection
* Email Risk Scoring
* Email Quarantine

---

## File Security Engine

Supported File Types:

* DOCM
* PDF
* ZIP
* PNG

Features:

* File Upload
* Macro Analysis
* PDF Analysis
* ZIP Analysis
* QR Analysis
* URL Extraction
* Threat Intelligence Checks

---

## Detection Engine

Features:

* YARA Rule Matching
* Sandbox Behavior Analysis
* Macro Detection
* PowerShell Detection
* AutoOpen Detection

---

## Threat Intelligence

Features:

* VirusTotal Hash Reputation
* URL Reputation Analysis
* Threat Intelligence Correlation
* SHA256 Hashing

---

## MITRE ATT&CK Mapping

Mapped Techniques:

* T1566.001 – Spearphishing Attachment
* T1566.002 – Spearphishing Link
* T1059.001 – PowerShell

---

## Risk Engine

Risk Levels:

* LOW
* MEDIUM
* HIGH
* CRITICAL

Verdicts:

* ALLOW
* SUSPICIOUS
* QUARANTINE

---

## Case Management

Features:

* Create Case
* View Cases
* Update Case
* Close Case
* Analyst Assignment
* Investigation Notes Timeline

---

## Search Engine

Features:

* Search Findings
* Search Files
* Search Analysis Results

---

## IOC Management

Features:

* IOC Export
* CSV Export

---

## Real-Time Alerting

Features:

* WebSocket Alerts
* Live Dashboard Updates

---

# 3. Dashboards

## SOC Dashboard

Metrics:

* Total Analyzed
* Allow
* Suspicious
* Quarantine
* Critical

---

## Threat Intelligence Dashboard

Features:

* Top Risk Files
* Threat Overview

---

## MITRE Dashboard

Features:

* Technique Frequency
* ATT&CK Statistics

---

## Incident Dashboard

Metrics:

* Total Incidents
* Open Incidents
* Closed Incidents
* Recent Incidents

---

# 4. Database Tables

users

* id
* username
* password

analysis_results

* id
* file_name
* risk_score
* risk_level
* verdict
* findings
* sha256
* urls
* mitre

cases

* id
* file_name
* analyst
* status
* notes
* created_at

case_notes

* id
* case_id
* analyst
* note
* created_at

---

# 5. APIs

Authentication

* POST /api/register
* POST /api/login

Analysis

* POST /api/analyze-email
* POST /api/analyze-file

Investigation

* GET /api/recent-findings
* GET /api/file/{file}

Cases

* POST /api/case
* GET /api/cases
* PUT /api/case-details/{id}
* POST /api/case-note
* GET /api/case-notes/{id}

Dashboards

* GET /api/dashboard
* GET /api/incident-dashboard
* GET /api/mitre-stats

Threat Intelligence

* GET /api/search
* GET /api/export/iocs

WebSocket

* GET /ws

---

# 6. Future Enhancements

* Docker Deployment
* Password Hashing (bcrypt)
* RBAC (Admin/Analyst)
* Audit Logs
* CI/CD Pipeline
* Kubernetes Deployment
* Automated Testing
* Hybrid Analysis Integration
* SOAR Automation
