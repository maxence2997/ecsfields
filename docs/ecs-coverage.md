# ECS coverage — v0.1.0

Pinned to **ECS 8.17**. This document tracks which ECS field families are
covered, deferred, or out of scope.

## Included

| Fieldset                                | Helpers  | Notes                                                                                                                                 |
| --------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `labels.*`, `tags`, `numeric_labels.*`  | 3        | Generic escape hatches: `Label`, `NumericLabel`, `Tags`                                                                               |
| `service.*`                             | 11       | `ServiceNodeRoles` is variadic                                                                                                        |
| `host.*` (top-level)                    | 9        | `HostIP` / `HostMAC` are variadic; `HostUptime` emits seconds                                                                         |
| `process.*` (top-level)                 | 11       | `ProcessUptime` emits seconds; `ProcessStart` is `time.Time`; endpoint-security subtrees excluded                                     |
| `event.*`                               | 22       | `event.duration` emits **nanoseconds**; `event.original` is bytes; typed enums for `kind`/`outcome`/`category`/`type`                 |
| `error.*` + `Err()`                     | 5 + 1    | `Err()` extracts `error.message` / `error.type` always, `error.stack_trace` when source implements `interface{ StackTrace() []byte }` |
| `log.*`                                 | 6        | Includes `log.origin.*`                                                                                                               |
| `trace.id`, `span.id`, `transaction.id` | 3        | APM correlation                                                                                                                       |
| `http.*`                                | 13       | Bytes are `int64`, status code is `int`                                                                                               |
| `url.*`                                 | 14       | `URLPort` is the only numeric field                                                                                                   |
| `client.*` (top-level)                  | 8        | Excludes network-monitoring subtrees                                                                                                  |
| `server.*` (top-level)                  | 8        | Mirrors `client.*`                                                                                                                    |
| `user_agent.*`                          | 4        | `original`, `name`, `version`, `device.name`                                                                                          |
| **Total**                               | **~115** |                                                                                                                                       |

## Deferred (additive in future v1.x)

These are valid ECS fields the library will eventually cover. They were left
out of v0.1.0 to ship a tight, useful surface first.

- `host.os.*`, `host.geo.*`
- `client.geo.*`, `client.as.*`, `server.geo.*`, `server.as.*`
- `log.syslog.*`
- `network.*`, `cloud.*`, `container.*`, `kubernetes.*`, `orchestrator.*`
- `user.*`

## Out of scope (not planned)

These fields target domains the library does not aim to serve.

- Runtime host metrics: `host.cpu.*`, `host.disk.*`, `host.network.*` (bandwidth)
- Security domain: `host.risk.*`, `event.risk_*`, `event.agent_id_status`
- NIDS: `client.bytes`, `client.packets`, `client.nat.*`, and `server.*` equivalents
- Endpoint security: `process.io.*`, `process.env_vars`, `process.vpid`,
  `process.entity_id`, `process.parent.*`, `process.group_leader.*`,
  `process.entry_leader.*`
- Specialized telemetry: `dns.*`, `tls.*`, `email.*`, `file.*`, `registry.*`
- Threat intel: `threat.*`, `vulnerability.*`, `code_signature.*`, `package.*`
