# chainwatch

`chainwatch` explores blockchain tooling with a small Go codebase and local fixtures. The technical goal is to detect reorgs, finality thresholds, and watcher alerts.

## Use Case

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Chainwatch Review Notes

`edge` and `stress` are the cases worth reading first. They show the optimistic and cautious ends of the fixture.

## Highlights

- `fixtures/domain_review.csv` adds cases for event finality and nonce pressure.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/chainwatch-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `settlement risk` and `nonce pressure`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Code Layout

The implementation keeps the scoring rule plain: reward signal and confidence, preserve slack, penalize drag, then classify the result into a review lane.

The Go addition stays small enough to inspect in one sitting.

## Run The Check

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Regression Path

The check exercises the source code and the review fixture. `edge` is the high score at 191; `stress` is the low score at 109.

## Future Work

The repository is intentionally scoped to local checks. I would expand it by adding adversarial fixtures before adding features.
