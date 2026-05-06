# ecsfields

Typed, ECS-compliant field constructors for structured logging in Go.

> Status: **beta** — API stabilizing.

ecsfields makes [Elastic Common Schema](https://www.elastic.co/guide/en/ecs/current/index.html)
non-compliance impossible by construction. Every constructor is hand-written, one-line,
and pinned to **ECS 8.17**. Currently ships a single sub-package — `zap` — that returns
`zap.Field` values with the correct ECS key and value type.

## Why ecsfields

Without ecsfields, ECS-compliant zap logs look like this:

```go
logger.Info("user signed in",
    zap.String("service.name", "auth-api"),
    zap.String("service.version", "1.2.3"),     // typo? compiles fine. silently dropped at ingest.
    zap.String("event.action", "user.login"),
    zap.String("event.outcome", "success"),      // wrong enum value? compiles fine.
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

Covers ~117 helpers across these top-level ECS fieldsets:

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

## Related projects

ecsfields composes with [`go.elastic.co/ecszap`](https://github.com/elastic/ecs-logging-go-zap)
(zap encoder for the log envelope).

For positioning detail — including ECS / Elasticsearch compatibility, when to
prefer `ecsf.Err(err)` over `zap.Error(err)`, and the full comparison table —
see [`docs/related-projects.md`](docs/related-projects.md).

## License

[MIT](LICENSE)
