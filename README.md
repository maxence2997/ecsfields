# ecsfields

Typed, ECS-compliant field constructors for structured logging in Go.

> Status: **beta** — API stabilizing toward v1.0.0.

ecsfields makes [Elastic Common Schema](https://www.elastic.co/guide/en/ecs/current/index.html)
non-compliance impossible by construction. Every constructor is hand-written, one-line,
and pinned to **ECS 8.17**. v0.1.0 ships a single sub-package — `zap` — that returns
`zap.Field` values with the correct ECS key and value type.

## Why ecsfields

Without ecsfields, ECS-compliant zap logs look like this:

```go
logger.Info("user signed in",
    zap.String("service.name", "auth-api"),
    zap.String("serivce.version", "1.2.3"),     // typo? compiles fine. silently dropped at ingest.
    zap.String("event.action", "user.login"),
    zap.String("event.outcome", "succes"),      // wrong enum value? compiles fine.
    zap.Int64("event.duration", int64(d)),      // ECS wants nanoseconds — easy to pass ms by mistake.
)
```

With ecsfields:

```go
import ecsf "github.com/maxence2997/ecsfields/zap"

logger.Info("user signed in",
    ecsf.ServiceName("auth-api"),
    ecsf.ServiceVersion("1.2.3"),                  // typed key — typos become compile errors.
    ecsf.EventAction("user.login"),
    ecsf.EventOutcome(ecsf.EventOutcomeSuccess),   // typed enum — compile-checked.
    ecsf.EventDuration(d),                         // takes time.Duration; emits ns automatically.
)
```

Every key is a typed function. Every value goes through a constructor that knows
the ECS-required type. Typos, wrong enum values, and unit mismatches become
compile errors instead of silent ingestion failures.

## Install

```bash
go get github.com/maxence2997/ecsfields/zap
```

## Quick start

Recommended setup pairs ecszap (envelope) with ecsfields (user fields):

```go
import (
    "os"

    "go.elastic.co/ecszap"
    "go.uber.org/zap"

    ecsf "github.com/maxence2997/ecsfields/zap"
)

encoderConfig := ecszap.NewDefaultEncoderConfig()
core := ecszap.NewCore(encoderConfig, os.Stdout, zap.InfoLevel)
logger := zap.New(core, zap.AddCaller())

logger.Info("user signed in",
    ecsf.ServiceName("auth-api"),
    ecsf.EventAction("user.login"),
    ecsf.EventOutcome(ecsf.EventOutcomeSuccess),
    ecsf.URLPath("/v1/login"),
    ecsf.HTTPRequestMethod("POST"),
    ecsf.HTTPResponseStatusCode(200),
)
```

Don't want ecszap? Replace the encoder/core/logger lines above with
`logger, _ := zap.NewProduction()`. ecsfields still works on its own — only
the envelope keys (`level`, `ts`, `msg`) will fall back to zap's defaults
instead of ECS-compliant names.

See [`example/main.go`](example/main.go) for a runnable end-to-end example.

## Coverage

v0.2.0 covers ~117 helpers across these top-level ECS fieldsets:

| Fieldset                             | Helpers            | Notes                                         |
| ------------------------------------ | ------------------ | --------------------------------------------- |
| `service.*`                          | 11                 | full top-level                                |
| `host.*`                             | 9                  | top-level only; metrics deferred              |
| `process.*`                          | 11                 | top-level; io / env_vars / entity_id deferred |
| `event.*`                            | 22 + 4 typed enums | kind / outcome / category / type are typed    |
| `error.*`                            | 5 + `Err()` / `ErrAny()` | full top-level                          |
| `log.*` (+ origin)                   | 6                  | syslog deferred                               |
| `trace` / `span` / `transaction` ids | 3                  | full                                          |
| `http.*`                             | 13                 | full                                          |
| `url.*`                              | 15                 | full                                          |
| `client.*`                           | 8                  | top-level; NIDS / nat deferred                |
| `server.*`                           | 8                  | top-level; NIDS / nat deferred                |
| `user_agent.*`                       | 4                  | full                                          |
| Generic escape hatches               | 3                  | `Label`, `NumericLabel`, `Tags`               |

See [`docs/ecs-coverage.md`](docs/ecs-coverage.md) for the full included / deferred /
out-of-scope matrix.

## ECS schema and Elasticsearch compatibility

ecsfields emits **ECS 8.17** field keys and types regardless of which Elasticsearch
or Kibana version you run. Older clusters will accept the documents via schema-on-write,
but Kibana dashboards designed for older ECS versions will not recognize 8.x-only
fields (`numeric_labels.*`, `service.address`, `service.ephemeral_id`, etc.).

## Relationship to other libraries

| Library                                                                 | What it does                                                                                                                                 | How it relates to ecsfields                                                                                                                                                                                                                                                                                                      |
| ----------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [`go.elastic.co/ecszap`](https://github.com/elastic/ecs-logging-go-zap) | zap encoder. Renames zap's built-in keys (`level` → `log.level`, `msg` → `message`, `ts` → `@timestamp`) and formats stack traces ECS-style. | **Use both — they cover different parts of the log.** ecszap fixes what zap auto-emits (`level`, `ts`, `msg`, `caller`, `stacktrace`). ecsfields gives you typed constructors for every field you add yourself. |
| [`github.com/elastic/ecs`](https://github.com/elastic/ecs)              | Canonical ECS schema definitions and document marshaling for full-document construction.                                                     | Different concern. ecsfields is for incremental field-by-field logging in zap, not for building complete ECS documents.                                                                                                                                                                                                          |
| [`github.com/andrewkroh/go-ecs`](https://github.com/andrewkroh/go-ecs)  | ECS schema query and introspection tool.                                                                                                     | Different concern (schema introspection at runtime, not log emission).                                                                                                                                                                                                                                                           |

### Using ecsfields error helpers vs `zap.Error`

Under the ecszap encoder, `zap.Error(err)` produces ECS `error.message`
(always) and `error.stack_trace` (only when `err` implements pkg/errors'
`StackTrace() errors.StackTrace`). Plain `errors.New(...)` does not produce
a stack. You do not have to migrate existing `zap.Error(err)` call sites
purely for ECS compliance.

That said, `ecsf.Err(err)` does three things `zap.Error(err)` doesn't:

1. **Always emits `error.type`.** Lets you filter Kibana by Go error class
   (`*pq.Error`, `*net.OpError`, ...). `zap.Error` never sets this field.
2. **Encoder-agnostic.** Produces correct ECS keys regardless of encoder.
   `zap.Error` only outputs ECS shape under ecszap; with the default JSON
   encoder it falls back to a flat `"error":"..."` string.
3. **Composable.** Single-field helpers (`ErrorCode`, `ErrorID`,
   `ErrorStackTrace`) and `ErrAny(any)` (for `recover()` values) cover cases
   that have no Go `error` to pass to `zap.Error` in the first place.

`ecsf.Err(err)` extracts `error.stack_trace` via either pkg/errors'
`StackTrace() errors.StackTrace` or samber/oops' `StackTrace() []byte`,
so any error wrapped by either library carries its stack through.

Recommended: use `ecsf.Err(err)` for new code or any error you want to
classify in Kibana. Keep existing `zap.Error(err)` call sites if you only
need message + stack_trace and you're committed to the ecszap encoder.

## License

[MIT](LICENSE)
