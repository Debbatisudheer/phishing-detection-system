# 12 — Security Review and Limitations

## Credential exposure

The archive contains a `.env` with a VirusTotal API credential. Treat it as exposed. Revoke/rotate it and remove credentials from Git history if necessary.

## Authentication

The source uses JWT, but the signing secret is hard-coded and passwords are handled directly rather than with a password hash. Production should use secret management and bcrypt/Argon2id.

## Authorization

Some endpoints have JWT middleware while others are directly registered. Authorization should be reviewed endpoint by endpoint. The role claim alone is not complete RBAC.

## WebSocket

The WebSocket upgrader accepts all origins. Restrict it to trusted frontend origins in production.

## Uploads

Uploaded filenames are used when constructing paths. Production should add filename normalization, traversal protection, type/size validation, random storage names and quarantine storage.

## SMTP

The current SMTP server allows insecure authentication and should be limited to controlled development/testing.

## Sandbox

Docker analysis is useful but should not automatically be treated as a complete malware-analysis security boundary.

## Detection limitations

- SPF/DKIM/DMARC is simulated.
- Threat intelligence includes static/demo lists.
- MITRE mapping is rule-based and incomplete.
- Risk scoring is deterministic and needs calibration.
- Some behavior/detonation findings are inferred rather than measured from real execution.

## Production roadmap

1. rotate secrets
2. hash passwords
3. externalize JWT secret
4. enforce authentication/RBAC
5. harden uploads
6. restrict WebSocket origins
7. secure SMTP
8. strengthen sandbox isolation
9. implement real email-authentication validation
10. improve intelligence lifecycle
11. add structured audit/observability
