# Changelog

## [Unreleased]

## [0.2.0]

### Added

- `ErrAny(v any) []zap.Field` — extracts ECS `error.*` fields from any value,
  intended for `recover()` payloads (typed as `any`). Delegates to `Err(err)`
  when the value satisfies `error`; falls back to `fmt.Sprint` /
  `fmt.Sprintf("%T", v)` for non-error values; returns nil for nil input.
- `Err(err)` now extracts `error.stack_trace` from `github.com/pkg/errors`
  errors as well — previously only `samber/oops`-style `StackTrace() []byte`
  was supported.

### Changed

- `Err(err)` stack-trace extraction is now delegated to an internal
  `extractStackTrace` helper that walks the error chain via `errors.As` and
  tries the `samber/oops` interface first, then the `pkg/errors` interface.
  No public API change.
- `ErrAny(v)` guards against typed-nil error inputs (e.g. `(*MyErr)(nil)`
  cast to `error`) — emits `error.message="<nil>"` + `error.type` instead of
  panicking inside `err.Error()`.

### Dependencies

- Add `github.com/pkg/errors v0.9.1` (direct) — required to type-check the
  pkg/errors `StackTrace() errors.StackTrace` interface.

### Documentation

- README clarifies the relationship between `ecsf.Err`, `ecsf.ErrorXxx`
  single-field helpers, and `zap.Error(err)` under the ecszap encoder; in
  particular that `zap.Error` only produces `error.stack_trace` when the
  underlying error implements pkg/errors' `StackTracer`.
