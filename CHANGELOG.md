# Changelog

## [Unreleased]

## [0.2.0]

### Added

- `ErrAny(v any) []zap.Field` — extracts ECS `error.*` fields from any value,
  intended for `recover()` payloads (typed as `any`). Delegates to `Err(err)`
  when the value satisfies `error`; falls back to `fmt.Sprint` /
  `fmt.Sprintf("%T", v)` for non-error values; returns nil for nil input.

### Documentation

- README clarifies the relationship between `ecsf.Err`, `ecsf.ErrorXxx`
  single-field helpers, and `zap.Error(err)` under the ecszap encoder.
