# Related projects

This document covers ECS / Elasticsearch compatibility, how ecsfields composes
with `go.elastic.co/ecszap`, and the most common positioning question — when
to reach for `ecsf.Err(err)` vs the standard `zap.Error(err)`.

## ECS schema and Elasticsearch compatibility

ecsfields emits **ECS 8.17** field keys and types regardless of which
Elasticsearch or Kibana version you run. Older clusters accept the documents
via schema-on-write — no rejection at ingest. The catch is dashboards: Kibana
visualizations designed for older ECS versions will not recognise 8.x-only
fields like `numeric_labels.*`, `service.address`, or `service.ephemeral_id`.

If you target a fixed older ECS version (e.g. ECS 1.7 with Elasticsearch
7.10), check the [ECS 8.17 reference](https://www.elastic.co/guide/en/ecs/8.17/index.html)
before relying on a field that may not exist in your dashboards.

## Composing with `go.elastic.co/ecszap`

ecszap is a zap encoder maintained by Elastic. It renames zap's built-in keys
to ECS equivalents (`level` → `log.level`, `msg` → `message`,
`ts` → `@timestamp`) and formats stack traces ECS-style. Covers roughly six
built-in keys that zap auto-emits per log line.

ecsfields and ecszap are complementary, non-overlapping:

- **ecszap** owns the *envelope* (the keys zap auto-emits regardless of what
  fields you pass).
- **ecsfields** owns the *user fields* (every `zap.Field` you explicitly add
  via `logger.Info(..., field, field, ...)`).

ecszap will not catch a typo like `zap.String("serivce.name", "x")` — it only
touches its six envelope keys. ecsfields gives you `ecsf.ServiceName("x")`,
which is a typed function call the compiler checks.

ecsfields does not require ecszap and works with any zap encoder.

## Using ecsfields error helpers vs `zap.Error`

This is the single most common positioning question, so it gets its own
section.

### TL;DR

| Goal                                             | Use                       |
| ------------------------------------------------ | ------------------------- |
| Log a Go `error` with full ECS shape             | `ecsf.Err(err)`        |
| Log a `recover()` payload (typed `any`)          | `ecsf.ErrAny(v)`          |
| Just need message + stack and already on ecszap  | `zap.Error(err)` is fine  |
| Need `error.type` populated for Kibana filtering | `ecsf.Err(err)`        |
| Run without ecszap (e.g. console encoder)        | `ecsf.Err(err)`        |

### Detailed comparison

Under the **ecszap encoder**, `zap.Error(err)` produces:

- `error.message` — always.
- `error.stack_trace` — only when `err` implements pkg/errors'
  `StackTrace() pkgerrors.StackTrace`. Plain `errors.New("...")` produces
  no stack.

Under any **other encoder** (default JSON, console, custom), `zap.Error(err)`
falls back to a flat string field `"error":"..."` — not ECS-compliant.

`ecsf.Err(err)` produces:

- `error.message` — always.
- `error.type` — always (e.g. `*pq.Error`, `*net.OpError`).
- `error.stack_trace` — when `err` implements either
  `StackTrace() []byte` (samber/oops) or
  `StackTrace() pkgerrors.StackTrace` (github.com/pkg/errors).

Output shape and key set are identical regardless of which encoder you use.

### Three things `ecsf.Err(err)` does that `zap.Error(err)` doesn't

1. **Always emits `error.type`.** This lets you build Kibana queries like
   `error.type:"*pq.Error"`. `zap.Error` never populates this field.
2. **Encoder-agnostic.** Produces ECS shape regardless of encoder. `zap.Error`
   only produces ECS shape under ecszap.
3. **Composable.** Single-field helpers (`ecsf.ErrorCode`, `ecsf.ErrorID`,
   `ecsf.ErrorStackTrace`) plus `ecsf.ErrAny(any)` (for `recover()` payloads)
   cover cases where there is no Go `error` value to pass to `zap.Error` in
   the first place.

### Recommendation

For new code or any error you want to classify in Kibana, use `ecsf.Err(err)`.
You do not need to migrate existing `zap.Error(err)` call sites purely for
ECS compliance — they already produce valid output if you are committed to
the ecszap encoder and only need message + (sometimes) stack.
