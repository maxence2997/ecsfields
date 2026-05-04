# ecsfields

Typed, ECS-compliant field constructors for structured logging in Go.

> Status: **beta** — API stabilizing toward v1.0.0.

ecsfields makes [Elastic Common Schema](https://www.elastic.co/guide/en/ecs/current/index.html)
non-compliance impossible by construction. Every constructor is hand-written, one-line,
and pinned to **ECS 8.17**. v0.1.0 ships a single sub-package — `zap` — that returns
`zap.Field` values with the correct ECS key and value type.

## Install

```bash
go get github.com/maxence2997/ecsfields/zap
```

## Quick start

```go
import (
    "go.uber.org/zap"

    ecszap "github.com/maxence2997/ecsfields/zap"
)

logger, _ := zap.NewProduction()
logger.Info("user signed in",
    ecszap.ServiceName("auth-api"),
    ecszap.EventAction("user.login"),
    ecszap.EventOutcome(ecszap.EventOutcomeSuccess),
    ecszap.URLPath("/v1/login"),
    ecszap.HTTPRequestMethod("POST"),
    ecszap.HTTPResponseStatusCode(200),
)
```

See [`example/main.go`](example/main.go) for a runnable end-to-end example.

## Coverage

v0.1.0 covers ~115 helpers across these top-level ECS fieldsets:

| Fieldset                             | Helpers            | Notes                                         |
| ------------------------------------ | ------------------ | --------------------------------------------- |
| `service.*`                          | 11                 | full top-level                                |
| `host.*`                             | 9                  | top-level only; metrics deferred              |
| `process.*`                          | 11                 | top-level; io / env_vars / entity_id deferred |
| `event.*`                            | 22 + 4 typed enums | kind / outcome / category / type are typed    |
| `error.*`                            | 5 + `Err()` helper | full top-level                                |
| `log.*` (+ origin)                   | 6                  | syslog deferred                               |
| `trace` / `span` / `transaction` ids | 3                  | full                                          |
| `http.*`                             | 13                 | full                                          |
| `url.*`                              | 14                 | full                                          |
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

| Library                                                                 | Role                                           | ecsfields relationship                                                                                                                       |
| ----------------------------------------------------------------------- | ---------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| [`go.elastic.co/ecszap`](https://github.com/elastic/ecs-logging-go-zap) | zap encoder that maps standard zap keys to ECS | Complementary. Use `ecszap.NewProductionEncoderConfig` for the encoder; use ecsfields for the field keys. ecsfields does not require ecszap. |
| [`github.com/elastic/ecs`](https://github.com/elastic/ecs)              | ECS document marshaling                        | Different concern (full document construction).                                                                                              |
| [`github.com/andrewkroh/go-ecs`](https://github.com/andrewkroh/go-ecs)  | ECS schema query                               | Different concern (schema introspection).                                                                                                    |

## License

[MIT](LICENSE)
