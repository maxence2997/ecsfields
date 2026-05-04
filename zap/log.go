// log.go — ECS log.* fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-log.html
//
// log.level / log.logger / log.original are usually set by the underlying log
// framework (zap). These helpers exist for callers manually constructing ECS
// payloads (e.g. log replay, sidecars).
//
// Out of scope (per spec): log.syslog.* (specialized syslog ingest).

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogLevel emits ECS log.level.
func LogLevel(level string) zapcore.Field { return zap.String("log.level", level) }

// LogLogger emits ECS log.logger (the named logger that produced the entry).
func LogLogger(name string) zapcore.Field { return zap.String("log.logger", name) }

// LogOriginal emits ECS log.original (raw bytes of the source log line).
func LogOriginal(b []byte) zapcore.Field { return zap.ByteString("log.original", b) }

// LogOriginFunction emits ECS log.origin.function.
func LogOriginFunction(fn string) zapcore.Field { return zap.String("log.origin.function", fn) }

// LogOriginFileName emits ECS log.origin.file.name.
func LogOriginFileName(name string) zapcore.Field {
	return zap.String("log.origin.file.name", name)
}

// LogOriginFileLine emits ECS log.origin.file.line.
func LogOriginFileLine(line int) zapcore.Field {
	return zap.Int64("log.origin.file.line", int64(line))
}
