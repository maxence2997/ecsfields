// trace.go — ECS trace.id / span.id / transaction.id top-level fields.
//
// ECS reference (8.17):
//   https://www.elastic.co/guide/en/ecs/8.17/ecs-tracing.html

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TraceID emits ECS trace.id.
func TraceID(id string) zapcore.Field { return zap.String("trace.id", id) }

// SpanID emits ECS span.id.
func SpanID(id string) zapcore.Field { return zap.String("span.id", id) }

// TransactionID emits ECS transaction.id.
func TransactionID(id string) zapcore.Field { return zap.String("transaction.id", id) }
