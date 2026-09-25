# 10 — Testing, QA and CI/CD

## Repository snapshot

The uploaded archive contains approximately 287 Go files, 112 Go test files, 55 QA Go tests and 50 frontend source JS/JSX files.

## QA areas

```text
qa/api
qa/e2e
qa/integration
qa/smoke
qa/load
qa/performance
qa/ui
qa/regression
```

API tests cover authentication, analysis, alerts, cases, campaigns, dashboards, incidents, investigation, IOC export, reports, sandbox, search, system health, threat intelligence and threat hunting.

Integration tests cover parser, attachment/macro, domain, email, ZIP, PDF, sandbox, Sigma, Splunk, STIX, threat feeds, threat intelligence, timeline and UEBA-related flows.

The archive also contains Playwright configuration, page objects, reports, traces, screenshots and performance scripts.

## Jenkins

Active stages:

1. Checkout
2. Environment check
3. Verify database
4. Import schema
5. Backend build
6. Frontend build
7. Backend tests

The Playwright Jenkins stage exists but is commented out.

## Makefile

```text
make fmt
make vet
make build
make backend-test
make frontend-build
make ci
make all
```

The presence of tests or old reports does not prove every test passes today. Fresh execution is required for a current pass/fail claim.
