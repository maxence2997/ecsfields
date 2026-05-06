// error.go — ECS error.* fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-error.html
//
// The package exposes single-field constructors (uniform with the rest of the
// library) plus one multi-field convenience helper, Err(), which extracts ECS
// error.* fields from a Go error without requiring any specific zap encoder.

package zap

import (
	"errors"
	"fmt"
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ErrorMessage emits ECS error.message.
func ErrorMessage(msg string) zapcore.Field { return zap.String("error.message", msg) }

// ErrorType emits ECS error.type (typically the Go type name of the error).
func ErrorType(t string) zapcore.Field { return zap.String("error.type", t) }

// ErrorStackTrace emits ECS error.stack_trace from raw bytes (e.g. the output
// of debug.Stack at a panic recovery site).
func ErrorStackTrace(b []byte) zapcore.Field { return zap.ByteString("error.stack_trace", b) }

// ErrorCode emits ECS error.code (a domain-specific identifier).
func ErrorCode(code string) zapcore.Field { return zap.String("error.code", code) }

// ErrorID emits ECS error.id (a unique identifier for the error instance).
func ErrorID(id string) zapcore.Field { return zap.String("error.id", id) }

// stackTracer is the conventional interface for errors that carry a captured stack.
// samber/oops satisfies this interface natively.
type stackTracer interface {
	StackTrace() []byte
}

// Err extracts ECS error.* fields from a Go error. It returns:
//
//   - error.message: always (err.Error())
//   - error.type:    always (fmt.Sprintf("%T", err))
//   - error.stack_trace: if any error in the chain implements interface{ StackTrace() []byte }
//
// The StackTrace method must have signature: StackTrace() []byte.
// Note: github.com/pkg/errors exposes StackTrace() errors.StackTrace ([]uintptr)
// and does NOT satisfy this interface. Use a wrapper or samber/oops instead.
//
// Err is the only multi-field constructor in the library, provided so callers
// do not need any specific zap encoder (e.g. ecszap) to obtain a stack trace.
// Returns nil if err is nil so the caller can splat it unconditionally.
func Err(err error) []zapcore.Field {
	if err == nil {
		return nil
	}
	fields := []zapcore.Field{
		ErrorMessage(err.Error()),
		ErrorType(fmt.Sprintf("%T", err)),
	}
	var st stackTracer
	if errors.As(err, &st) {
		if stack := st.StackTrace(); len(stack) > 0 {
			fields = append(fields, ErrorStackTrace(stack))
		}
	}
	return fields
}

// ErrAny extracts ECS error.* fields from any value, intended for cases where
// the input is not statically typed as error — most commonly the result of
// recover() during panic handling. Behavior by input type:
//
//   - nil:             returns nil (no fields)
//   - typed-nil error: error.type emitted, error.message = "<nil>" — never
//     calls Error() on the typed-nil receiver, which would panic
//   - error:           delegates to Err(err) — error.stack_trace included if
//     the error implements interface{ StackTrace() []byte }
//   - other:           error.message = fmt.Sprint(v); error.type = fmt.Sprintf("%T", v)
//
// ErrAny intentionally does not call runtime/debug.Stack() itself. To attach
// the panic stack, append ErrorStackTrace(debug.Stack()) at the call site —
// callers may want to skip the cost or use a different stack source.
//
// Typical panic recovery:
//
//	defer func() {
//	    if r := recover(); r != nil {
//	        fields := ErrAny(r)
//	        fields = append(fields, ErrorStackTrace(debug.Stack()))
//	        logger.Error("panic recovered", fields...)
//	    }
//	}()
func ErrAny(v any) []zapcore.Field {
	if v == nil {
		return nil
	}
	if err, ok := v.(error); ok {
		if isTypedNil(err) {
			return []zapcore.Field{
				ErrorMessage("<nil>"),
				ErrorType(fmt.Sprintf("%T", err)),
			}
		}
		return Err(err)
	}
	return []zapcore.Field{
		ErrorMessage(fmt.Sprint(v)),
		ErrorType(fmt.Sprintf("%T", v)),
	}
}

// isTypedNil reports whether v is non-nil at the interface level but holds a
// nil concrete value (e.g. (*MyErr)(nil) cast to error). Calling methods that
// dereference the receiver on such a value panics, so ErrAny short-circuits
// before invoking err.Error().
func isTypedNil(v any) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func, reflect.Interface:
		return rv.IsNil()
	}
	return false
}
