# 06 — Risk Engine and Decision

The risk engine receives subject, body, URLs and findings.

It adds a URL contribution, configured finding weights, subject keywords and body keywords. The score is capped at 1000.

The decision engine currently uses:

```text
<100       ALLOW
100–399    SUSPICIOUS
>=400      QUARANTINE
```

Risk level and verdict are different:

- risk level = how dangerous the evidence appears
- verdict = what the platform decides to do

The current model is deterministic and rule-based, not an ML risk model. Its main advantage is explainability: analysts can inspect the findings that contributed to the result.
