# AGENTS.md

## Code review

Report findings first: the problem, the minimal fix, and where it goes. Wait for approval before editing code.

## Tests

Ultra-minimal: a handful of focused cases, properties and brute-force cross-checks preferred over fixtures, no boilerplate helpers. Keep them in a `#[cfg(test)] mod tests` in the same file.

## Bugs

When a test exposes a bug, leave the implementation alone and make the test red — `#[should_panic]` for a crash, a plain assertion for expected behaviour. The failing test is the report.
