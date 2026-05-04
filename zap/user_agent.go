// user_agent.go — ECS user_agent.* fields.
//
// ECS reference (8.17):
//   https://www.elastic.co/guide/en/ecs/8.17/ecs-user_agent.html

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// UserAgentOriginal emits ECS user_agent.original.
func UserAgentOriginal(v string) zapcore.Field { return zap.String("user_agent.original", v) }

// UserAgentName emits ECS user_agent.name.
func UserAgentName(v string) zapcore.Field { return zap.String("user_agent.name", v) }

// UserAgentVersion emits ECS user_agent.version.
func UserAgentVersion(v string) zapcore.Field { return zap.String("user_agent.version", v) }

// UserAgentDeviceName emits ECS user_agent.device.name.
func UserAgentDeviceName(v string) zapcore.Field {
	return zap.String("user_agent.device.name", v)
}
