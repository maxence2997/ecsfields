# CLAUDE.md — ecsfields

> This file is the **single source of truth** for all AI agents working on this project.
> Both `.github/copilot-instructions.md` and `AGENTS.md` reference this file.

## Project Overview

ecsfields is a small, hand-written Go library providing **typed, ECS-compliant field
constructors** for structured logging. Module path: `github.com/maxence2997/ecsfields`.

The library makes ECS-non-compliant log fields (typo'd keys, wrong value types)
impossible by construction. v0.1.0 ships zap-only support; slog and other loggers
will land as additive sibling sub-packages without breaking changes.

ECS schema version pinned: **8.17**.

## Package Layout

| Path       | Description                                             |
| ---------- | ------------------------------------------------------- |
| `.` (root) | Namespace package. `doc.go` only — no exported symbols. |
| `zap/`     | `zap.Field` constructors for ECS 8.17 fields.           |

Import the sub-package directly:

```go
import ecszap "github.com/maxence2997/ecsfields/zap"
```

## File Index (zap/)

| File                                   | Contents                                       |
| -------------------------------------- | ---------------------------------------------- |
| `doc.go`                               | Sub-package GoDoc                              |
| `service.go` / `service_test.go`       | `service.*` constructors                       |
| `host.go` / `host_test.go`             | `host.*` top-level constructors                |
| `process.go` / `process_test.go`       | `process.*` top-level constructors             |
| `event.go` / `event_test.go`           | `event.*` constructors and typed enums         |
| `error.go` / `error_test.go`           | `error.*` constructors and `Err()` helper      |
| `log.go` / `log_test.go`               | `log.*` and `log.origin.*` constructors        |
| `trace.go` / `trace_test.go`           | `trace.id`, `span.id`, `transaction.id`        |
| `http.go` / `http_test.go`             | `http.*` request and response fields           |
| `url.go` / `url_test.go`               | `url.*` constructors                           |
| `client.go` / `client_test.go`         | `client.*` top-level constructors              |
| `server.go` / `server_test.go`         | `server.*` top-level constructors              |
| `user_agent.go` / `user_agent_test.go` | `user_agent.*` constructors                    |
| `labels.go` / `labels_test.go`         | `Label`, `NumericLabel`, `Tags` escape hatches |

## Development Workflow

```bash
make help          # list all available make targets
make fmt           # gofmt + goimports
make lint          # go vet + golangci-lint
make test          # race-detector tests (default count=10)
make check         # fmt-check + lint + test (pre-commit gate)
make test-cover    # coverage report -> coverage.html
make tidy          # go mod tidy
make deps          # go mod download + tidy
make clean         # remove local coverage artifacts
```

## Conventions

- **Architecture**: SRP per constructor (one ECS field per func). Open-Closed via
  additive function calls and additive sub-packages. No interfaces — pure function
  library.
- **No implicit dependencies**: lib must work with any zap encoder. Do not require
  `ecszap` or any specific encoder.
- **Naming**: PascalCase ↔ dotted ECS key. `EventAction` → `event.action`.
  `ServiceNodeName` → `service.node.name`. Numeric → string variants suffixed
  with `Int` (e.g. `EventCodeInt`).
- **Native Go types in signatures**: `time.Duration` for durations, `int` for
  counts. Convert to ECS-required types inside the constructor.
- **Each field file** must reference upstream ECS YAML in a top comment.
- **Go style**: `gofmt`/`goimports`, GoDoc on all public symbols.
- **Markdown**: no emojis in documentation files.
- **Language consistency**: code and commit messages in English. Respond in the
  user's language (Traditional Chinese when applicable).
- **File encoding**: UTF-8 without BOM.

## Commit Messages

Follow `.github/instructions/commit-message-instructions.md`.

Format:

```
<type>: <subject>          (max 50 chars)

1.<reason> -> <change>     (max 50 chars each, max 3 items)
2.<reason> -> <change>
```

Types: `feat`, `fix`, `refactor`, `doc`, `style`, `test`, `chore`, `revert`, `merge`, `sync`.

Rules: one logical change per commit. English only. No half-word abbreviations.

## Git

- Branch strategy: `feat/<name>`, `refactor/<name>`, `bugfix/<name>`, `fix/<name>`, `chore/<name>`.
  Never push directly to `main`. Open a PR from a feature branch into `main`.
- Pull request description: follow `.github/PULL_REQUEST_TEMPLATE.md`. Fill in every section.

## Critical Rules

1. **Read before write** -- read the target file fully before any edit.
2. **Test first, fix second** -- TDD is the only allowed flow. Production code
   must follow a failing test. No exceptions.
3. **`make check` gates every commit** -- `fmt -> lint -> race-detector test` must pass.
   Skip only for documentation-only commits.
4. **Minimal changes** -- one concern per edit; no drive-by refactors.
5. **No breaking changes without version bump** -- never rename, remove, or change the
   signature of an exported symbol without a major version bump.
6. **Accuracy** -- do not make assumptions about ECS field types. Verify against the
   pinned ECS 8.17 spec.
7. **STOP -- before every commit, verify this checklist**:
   1. Run `make check` and confirm it passes.
   2. Commit message follows `commit-message-instructions.md`.
   3. This commit contains exactly one logical change.

## ECS Coverage

See `docs/ecs-coverage.md` for the in-scope / deferred / out-of-scope matrix.
