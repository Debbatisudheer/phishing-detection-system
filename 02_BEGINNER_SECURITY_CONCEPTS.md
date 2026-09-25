# 02 — Beginner Security Concepts

## Phishing

A phishing attack tries to make a person click a link, open an attachment, reveal credentials or perform another unsafe action.

## IOC

IOC = Indicator of Compromise. Examples: URL, domain, IP, email address and file hash.

## Hash

A SHA-256 hash is a fingerprint-like value for a file. The project compares hashes with local and VirusTotal reputation.

## Threat intelligence

Threat intelligence provides known information about suspicious infrastructure or files. This project uses VirusTotal, a PhishTank-style feed, local domains/IPs/hashes and a local URL reputation table.

## YARA

YARA is rule-based pattern matching. The supplied rules include patterns for PowerShell downloaders, encoded PowerShell and registry persistence.

## Sandbox

A sandbox is an isolated analysis environment. The project has a Docker path with resource limits and a timeout. Some other behavior analysis is simulated from content.

## Risk scoring

Findings contribute weights to a numeric score. The score is capped at 1000.

## Verdict

The decision engine currently maps:

```text
<100       ALLOW
100–399    SUSPICIOUS
>=400      QUARANTINE
```

## MITRE ATT&CK

The project maps selected findings to ATT&CK-style techniques such as spearphishing attachment/link and PowerShell.

## SOC

SOC = Security Operations Center. The platform supports alert → investigation → IOC correlation → case → notes → report.
