# Review Journal

The review surface for `chainwatch` is deliberately narrow: one fixture, one scoring rule, and one local check.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its blockchain tooling focus without claiming live deployment or external usage.

## Cases

- `baseline`: `event finality`, score 129, lane `watch`
- `stress`: `nonce pressure`, score 109, lane `watch`
- `edge`: `settlement risk`, score 191, lane `ship`
- `recovery`: `proof depth`, score 171, lane `ship`
- `stale`: `event finality`, score 185, lane `ship`

## Note

A future change should add new cases before it changes the scoring rule.
