// url.go — ECS url.* fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-url.html

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// URLOriginal emits ECS url.original.
func URLOriginal(v string) zapcore.Field { return zap.String("url.original", v) }

// URLFull emits ECS url.full.
func URLFull(v string) zapcore.Field { return zap.String("url.full", v) }

// URLScheme emits ECS url.scheme.
func URLScheme(v string) zapcore.Field { return zap.String("url.scheme", v) }

// URLDomain emits ECS url.domain.
func URLDomain(v string) zapcore.Field { return zap.String("url.domain", v) }

// URLRegisteredDomain emits ECS url.registered_domain.
func URLRegisteredDomain(v string) zapcore.Field { return zap.String("url.registered_domain", v) }

// URLTopLevelDomain emits ECS url.top_level_domain.
func URLTopLevelDomain(v string) zapcore.Field { return zap.String("url.top_level_domain", v) }

// URLSubdomain emits ECS url.subdomain.
func URLSubdomain(v string) zapcore.Field { return zap.String("url.subdomain", v) }

// URLPort emits ECS url.port.
func URLPort(port int) zapcore.Field { return zap.Int64("url.port", int64(port)) }

// URLPath emits ECS url.path.
func URLPath(v string) zapcore.Field { return zap.String("url.path", v) }

// URLQuery emits ECS url.query.
func URLQuery(v string) zapcore.Field { return zap.String("url.query", v) }

// URLFragment emits ECS url.fragment.
func URLFragment(v string) zapcore.Field { return zap.String("url.fragment", v) }

// URLUsername emits ECS url.username.
func URLUsername(v string) zapcore.Field { return zap.String("url.username", v) }

// URLPassword emits ECS url.password.
//
// WARNING: This field contains credentials. In production, redact the value
// before logging, or use URLPasswordRedacted to always emit "***".
func URLPassword(v string) zapcore.Field { return zap.String("url.password", v) }

// URLPasswordRedacted emits url.password as a fixed redaction marker ("***").
// Prefer this over URLPassword when the credential must not appear in logs.
func URLPasswordRedacted() zapcore.Field { return zap.String("url.password", "***") }

// URLExtension emits ECS url.extension.
func URLExtension(v string) zapcore.Field { return zap.String("url.extension", v) }
