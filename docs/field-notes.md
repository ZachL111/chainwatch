# Field Notes

The fixture is small on purpose, which makes each domain case carry real weight.

The domain cases cover `event finality`, `nonce pressure`, `settlement risk`, and `proof depth`. They sit beside the smaller starter fixture so the project has both a compact scoring check and a domain-flavored review check.

`edge` is the strongest case at 191 on `settlement risk`. `stress` is the cautious anchor at 109 on `nonce pressure`.

The language-specific addition keeps the review model as a small package with table tests.
