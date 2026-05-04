# AGENTS.md — ecsfields

> Canonical reference: [`CLAUDE.md`](CLAUDE.md). This file provides agent-specific guidance
> that supplements the shared rules defined there. Read `CLAUDE.md` first.

## For All Agents

- This is a Go library. Production dependency on `go.uber.org/zap` only.
  Do not add other external production dependencies without explicit user approval.
- `make check` (fmt + lint + race-detector test) must pass before any commit.
- One logical change per commit. Follow the commit message format in
  `.github/instructions/commit-message-instructions.md`.

## Code Agent

When writing or modifying code:

1. Read the target file in full before editing.
2. Every constructor maps to exactly one ECS field. Each file documents the
   upstream ECS YAML reference in a top comment.
3. Do not introduce runtime polymorphism or interfaces. The library is a pure
   function library by design (D5 in spec).
4. Do not couple to a specific zap encoder (no `ecszap` dependency).
5. Run `make check` after every change to confirm nothing is broken.
6. Do not add, rename, or remove exported symbols without a major version bump discussion.

## Test Agent

When writing or modifying tests:

1. Use table-driven tests for groups of similar constructors.
2. Each helper test asserts: ECS key, zap field type, value.
3. Use `github.com/stretchr/testify/assert` and `require` consistently.
4. For typed enums, also assert the constant string values match the ECS spec.

## Documentation Agent

When editing documentation:

1. Keep `docs/ecs-coverage.md` in sync with what is exported in `zap/`.
2. No emojis in markdown files.
3. Keep `CHANGELOG.md` up to date for user-facing changes.

## Review Agent

When reviewing code or PRs:

1. Verify each constructor uses the correct zap value type for the ECS field type.
2. Verify duration helpers convert to the unit ECS specifies (e.g. nanoseconds
   for `event.duration`).
3. Verify `make check` passes on the branch under review.
