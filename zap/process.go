// process.go — ECS process.* top-level fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-process.html
//
// Out of scope (per spec, security-domain or endpoint-monitoring fields):
//   process.io.*, process.env_vars, process.entity_id, process.vpid,
//   process.parent.*, process.group_leader.*, process.entry_leader.*
//
// See docs/ecs-coverage.md for full scope decisions.

package zap

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ProcessArgs emits ECS process.args (keyword[]).
func ProcessArgs(args ...string) zapcore.Field { return zap.Strings("process.args", args) }

// ProcessArgsCount emits ECS process.args_count.
func ProcessArgsCount(n int) zapcore.Field {
	return zap.Int64("process.args_count", int64(n))
}

// ProcessCommandLine emits ECS process.command_line.
func ProcessCommandLine(v string) zapcore.Field { return zap.String("process.command_line", v) }

// ProcessExecutable emits ECS process.executable.
func ProcessExecutable(v string) zapcore.Field { return zap.String("process.executable", v) }

// ProcessExitCode emits ECS process.exit_code.
func ProcessExitCode(code int) zapcore.Field {
	return zap.Int64("process.exit_code", int64(code))
}

// ProcessName emits ECS process.name.
func ProcessName(v string) zapcore.Field { return zap.String("process.name", v) }

// ProcessPID emits ECS process.pid.
func ProcessPID(pid int) zapcore.Field { return zap.Int64("process.pid", int64(pid)) }

// ProcessStart emits ECS process.start as a date.
func ProcessStart(t time.Time) zapcore.Field { return zap.Time("process.start", t) }

// ProcessTitle emits ECS process.title.
func ProcessTitle(v string) zapcore.Field { return zap.String("process.title", v) }

// ProcessUptime emits ECS process.uptime as seconds.
func ProcessUptime(d time.Duration) zapcore.Field {
	return zap.Int64("process.uptime", int64(d.Seconds()))
}

// ProcessWorkingDirectory emits ECS process.working_directory.
func ProcessWorkingDirectory(v string) zapcore.Field {
	return zap.String("process.working_directory", v)
}
