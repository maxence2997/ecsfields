// host.go — ECS host.* top-level fields (excluding metrics).
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-host.html
//
// Out of scope (spec D5 / "deferred or not planned"):
//   host.cpu.*, host.disk.*, host.network.*  (runtime metrics)
//   host.os.*, host.geo.*                     (deferred to future v1.x)
//   host.risk.*                               (security domain, not planned)
//
// See docs/ecs-coverage.md for full scope decisions.

package zap

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HostArchitecture emits ECS host.architecture (e.g. "arm64", "x86_64").
func HostArchitecture(v string) zapcore.Field { return zap.String("host.architecture", v) }

// HostDomain emits ECS host.domain (Active Directory or DNS domain).
func HostDomain(v string) zapcore.Field { return zap.String("host.domain", v) }

// HostHostname emits ECS host.hostname (the canonical hostname, may include FQDN).
func HostHostname(v string) zapcore.Field { return zap.String("host.hostname", v) }

// HostID emits ECS host.id (a unique identifier for the host).
func HostID(v string) zapcore.Field { return zap.String("host.id", v) }

// HostIP emits ECS host.ip as an array of IP address strings.
func HostIP(ips ...string) zapcore.Field { return zap.Strings("host.ip", ips) }

// HostMAC emits ECS host.mac as a keyword array of MAC addresses.
func HostMAC(macs ...string) zapcore.Field { return zap.Strings("host.mac", macs) }

// HostName emits ECS host.name (the operator-assigned name of the host).
func HostName(v string) zapcore.Field { return zap.String("host.name", v) }

// HostType emits ECS host.type (e.g. "vm", "container", "physical").
func HostType(v string) zapcore.Field { return zap.String("host.type", v) }

// HostUptime emits ECS host.uptime as seconds since boot.
func HostUptime(d time.Duration) zapcore.Field {
	return zap.Int64("host.uptime", int64(d/time.Second))
}
