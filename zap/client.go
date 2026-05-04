// client.go — ECS client.* top-level fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-client.html
//
// Out of scope (per spec, network monitoring fields):
//   client.bytes, client.packets, client.nat.*, client.geo.*, client.as.*,
//   client.user.*

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ClientAddress emits ECS client.address.
func ClientAddress(v string) zapcore.Field { return zap.String("client.address", v) }

// ClientIP emits ECS client.ip.
func ClientIP(v string) zapcore.Field { return zap.String("client.ip", v) }

// ClientPort emits ECS client.port.
func ClientPort(port int) zapcore.Field { return zap.Int64("client.port", int64(port)) }

// ClientMAC emits ECS client.mac.
func ClientMAC(v string) zapcore.Field { return zap.String("client.mac", v) }

// ClientDomain emits ECS client.domain.
func ClientDomain(v string) zapcore.Field { return zap.String("client.domain", v) }

// ClientRegisteredDomain emits ECS client.registered_domain.
func ClientRegisteredDomain(v string) zapcore.Field {
	return zap.String("client.registered_domain", v)
}

// ClientTopLevelDomain emits ECS client.top_level_domain.
func ClientTopLevelDomain(v string) zapcore.Field {
	return zap.String("client.top_level_domain", v)
}

// ClientSubdomain emits ECS client.subdomain.
func ClientSubdomain(v string) zapcore.Field { return zap.String("client.subdomain", v) }
