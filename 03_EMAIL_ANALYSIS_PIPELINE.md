# 03 — Email Analysis Pipeline

## Input

Email enters through the SMTP server on port 2525 or `POST /api/analyze-email`.

## Parsing

`internal/parser/email_parser.go` extracts From, Reply-To, Return-Path, Subject, Body and attachments. Multipart attachments are saved under `uploads/`.

## Detection flow

```text
Email
 ↓
URLs / IOCs
 ↓
YARA + sandbox-style body analysis
 ↓
Header + display-name checks
 ↓
BEC + thread-hijack checks
 ↓
Email authentication simulation
 ↓
Sender reputation/history
 ↓
Campaign + QR + ZIP checks
 ↓
Domain + URL reputation + VirusTotal + feeds
 ↓
Attachment analysis
 ↓
Hash reputation
 ↓
Risk score
 ↓
Decision + MITRE
 ↓
DB + reports + WebSocket
```

## Important limitation

The SPF/DKIM/DMARC package is simulation logic based on sender domain; it is not a full standards-compliant verifier.

## Attachments

ZIP files can be extracted and checked for nested ZIPs. DOCM/XLSM files can have macro extraction and YARA/behavior analysis. PDFs have text and URL extraction. Images can be checked for QR codes.

## Outputs

The pipeline can save email results, IOC data, STIX/Sigma/report artifacts and a Splunk-style JSON event, and can broadcast a live WebSocket event.
