# Copilot Instructions -- ecsfields

> **Single source of truth**: [`CLAUDE.md`](../CLAUDE.md) at the project root.
> This file re-exports the shared rules for GitHub Copilot. All project conventions,
> development workflow, critical rules, and session protocol are defined there.
> When in doubt, follow `CLAUDE.md`.

## Quick Reference

- Module: `github.com/maxence2997/ecsfields` -- typed ECS-compliant field
  constructors for structured logging. v0.1.0 ships zap-only.
- ECS schema: pinned to 8.17.
- Sub-package: `github.com/maxence2997/ecsfields/zap`.
- Pre-commit gate: `make check` (fmt + lint + race-detector test).
- Commit format: `.github/instructions/commit-message-instructions.md`.
- PR template: `.github/PULL_REQUEST_TEMPLATE.md`.

## Copilot-Specific Notes

- Copilot Chat and Copilot Workspace should read `CLAUDE.md` for the full rule set.
- Inline completions: prefer `assert`/`require` from `testify` in test files.
- Do not generate code that adds external production dependencies beyond `go.uber.org/zap`.
- Each new helper must come with: GoDoc, a top-of-file ECS YAML reference,
  and a table-driven test covering key + type + value.
- Respond in the user's language (Traditional Chinese when applicable).
  Code and commit messages are always in English.
