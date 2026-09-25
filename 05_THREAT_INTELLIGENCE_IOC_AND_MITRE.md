# 05 — Threat Intelligence, IOC and MITRE

## IOC types

URLs, domains, IPs, hashes, email indicators and attachment references are handled.

## Intelligence sources

- VirusTotal integration
- local PhishTank-style feed
- local malicious domain/IP lists
- local malicious hash list
- PostgreSQL URL reputation
- threat-intelligence cache

## IOC correlation

`ioc_correlation` stores IOC/source/file relationships plus first seen, last seen and hit count. This supports campaign-style investigation.

## MITRE

The source maps selected findings to techniques including:

- T1566.001 — Spearphishing Attachment
- T1566.002 — Spearphishing Link
- T1059.001 — PowerShell

Mapping is rule-based and not complete ATT&CK coverage.

## STIX

The project creates a STIX-style URL indicator JSON.

## Sigma

The project generates a Sigma-style YAML rule from sender, risk score and MITRE information.

## Splunk

The project exports JSON containing timestamp, event type, sender, subject, risk score, decision and MITRE.

## Credential warning

The uploaded archive contains a VirusTotal API credential in `.env`. It is intentionally omitted from this documentation. Rotate/revoke it.
