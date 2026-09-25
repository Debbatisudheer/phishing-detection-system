# 04 — File Analysis and Sandbox

## Supported paths

- DOCM / XLSM macro analysis
- PS1 analysis
- ZIP and nested ZIP analysis
- PDF extraction/analysis
- PNG/JPG/JPEG QR analysis
- SHA-256 hashing
- VirusTotal hash lookup
- YARA
- sandbox jobs

## Upload flow

```text
Browser
 ↓
POST /api/analyze-file
 ↓
uploads/<filename>
 ↓
file-type analyzer
 ↓
risk + MITRE + verdict
```

## Docker sandbox

`internal/sandbox/docker_executor.go` runs an analysis container with `--memory=512m`, `--cpus=1`, a 120-second Go context timeout, mounted sample/rules and commands for file type, strings, SHA-256, ClamAV and YARA.

## Important distinction

The project has static/content analysis, rule-based behavior analysis and a Docker execution path. This is not the same as a complete commercial malware sandbox with VM isolation, kernel telemetry, memory forensics and full behavioral tracing.

## Safety

Use only controlled test samples. Do not execute unknown malware on a normal workstation.
