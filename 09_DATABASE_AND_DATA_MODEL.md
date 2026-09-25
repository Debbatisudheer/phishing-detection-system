# 09 — Database and Data Model

PostgreSQL is connected through Go `database/sql` and `lib/pq`.

The schema contains:

```text
alerts
analysis_results
analyst_notes
campaigns
case_notes
cases
docker_reports
emails
ioc_correlation
mitre_events
sandbox_jobs
sandbox_reports
threat_intel_cache
url_reputation
users
```

## Investigation lifecycle

```text
Email/file analyzed
 ↓
analysis result
 ↓
IOC correlation
 ↓
risk + verdict
 ↓
alert
 ↓
case
 ↓
case notes
 ↓
investigation/report
```

## Repository design

Database operations are organized into feature repositories such as users, analysis, cases, campaigns, IOC, sandbox, threat intelligence and incidents.

## CI schema

`database/schema.sql` is imported by Jenkins into the CI PostgreSQL database.

## Production improvements

Consider foreign keys, indexes, migrations, retention, backups, audit trails, encryption and stronger data-access controls.
