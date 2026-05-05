// server.go — ECS server.* top-level fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-server.html
//
// Out of scope (per spec, network monitoring fields):
//   server.bytes, server.packets, server.nat.*, server.geo.*, server.as.*,
//   server.user.*
//
// See docs/ecs-coverage.md for full scope decisions.

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ServerAddress emits ECS server.address.
func ServerAddress(v string) zapcore.Field { return zap.String("server.address", v) }

// ServerIP emits ECS server.ip.
func ServerIP(v string) zapcore.Field { return zap.String("server.ip", v) }

// ServerPort emits ECS server.port.
func ServerPort(port int) zapcore.Field { return zap.Int64("server.port", int64(port)) }

// ServerMAC emits ECS server.mac.
func ServerMAC(v string) zapcore.Field { return zap.String("server.mac", v) }

// ServerDomain emits ECS server.domain.
func ServerDomain(v string) zapcore.Field { return zap.String("server.domain", v) }

// ServerRegisteredDomain emits ECS server.registered_domain.
func ServerRegisteredDomain(v string) zapcore.Field {
	return zap.String("server.registered_domain", v)
}

// ServerTopLevelDomain emits ECS server.top_level_domain.
func ServerTopLevelDomain(v string) zapcore.Field {
	return zap.String("server.top_level_domain", v)
}

// ServerSubdomain emits ECS server.subdomain.
func ServerSubdomain(v string) zapcore.Field { return zap.String("server.subdomain", v) }
