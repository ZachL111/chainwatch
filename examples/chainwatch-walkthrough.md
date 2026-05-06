# Chainwatch Walkthrough

The fixture is intentionally compact, so the review starts with the cases that pull farthest apart.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | event finality | 129 | watch |
| stress | nonce pressure | 109 | watch |
| edge | settlement risk | 191 | ship |
| recovery | proof depth | 171 | ship |
| stale | event finality | 185 | ship |

Start with `edge` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `settlement risk` against `nonce pressure`, not the raw score alone.
