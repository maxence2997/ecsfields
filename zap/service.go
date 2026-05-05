// service.go — ECS service.* top-level fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-service.html
// Source CSV:           https://github.com/elastic/ecs/blob/v8.17.0/generated/csv/fields.csv
//
// All service.* fields are keyword in ECS, except service.node.roles which is
// a keyword array.

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ServiceAddress emits ECS service.address.
func ServiceAddress(v string) zapcore.Field { return zap.String("service.address", v) }

// ServiceEnvironment emits ECS service.environment.
func ServiceEnvironment(v string) zapcore.Field { return zap.String("service.environment", v) }

// ServiceEphemeralID emits ECS service.ephemeral_id.
func ServiceEphemeralID(v string) zapcore.Field { return zap.String("service.ephemeral_id", v) }

// ServiceID emits ECS service.id.
func ServiceID(v string) zapcore.Field { return zap.String("service.id", v) }

// ServiceName emits ECS service.name.
func ServiceName(v string) zapcore.Field { return zap.String("service.name", v) }

// ServiceNodeName emits ECS service.node.name.
func ServiceNodeName(v string) zapcore.Field { return zap.String("service.node.name", v) }

// ServiceNodeRole emits ECS service.node.role.
//
// Note: in ECS 8.x the multi-role variant is service.node.roles (keyword[]).
// service.node.role is retained for single-role services.
func ServiceNodeRole(v string) zapcore.Field { return zap.String("service.node.role", v) }

// ServiceNodeRoles emits ECS service.node.roles (keyword array).
func ServiceNodeRoles(roles ...string) zapcore.Field {
	return zap.Strings("service.node.roles", roles)
}

// ServiceState emits ECS service.state.
func ServiceState(v string) zapcore.Field { return zap.String("service.state", v) }

// ServiceType emits ECS service.type.
func ServiceType(v string) zapcore.Field { return zap.String("service.type", v) }

// ServiceVersion emits ECS service.version.
func ServiceVersion(v string) zapcore.Field { return zap.String("service.version", v) }
