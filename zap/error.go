// error.go — ECS error.* fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-error.html
//
// The package exposes single-field constructors (uniform with the rest of the
// library) plus two multi-field convenience helpers, Err() and ErrAny(), that
// pack the standard error.* fields into a single inline zap.Field. Callers do
// not need any specific zap encoder — output is flat dotted ECS keys, the
// same shape every other constructor in this package emits.

package zap

import (
	"errors"
	"fmt"
	"reflect"

	pkgerrors "github.com/pkg/errors"
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

// stackTracerBytes is satisfied by errors that expose a pre-formatted stack
// trace as bytes. samber/oops errors satisfy this interface natively.
type stackTracerBytes interface {
	StackTrace() []byte
}

// stackTracerPCs is satisfied by errors that expose a stack trace as
// pkg/errors-style program counters. github.com/pkg/errors errors (e.g. those
// returned by pkgerrors.New / pkgerrors.Wrap) satisfy this interface natively.
type stackTracerPCs interface {
	StackTrace() pkgerrors.StackTrace
}

// Err returns a single inline zap.Field that emits ECS error.* fields:
//
//   - error.message: always (err.Error())
//   - error.type:    always (fmt.Sprintf("%T", err))
//   - error.stack_trace: if any error in the chain implements one of the
//     conventional stack-trace interfaces — checked in this order:
//     1. interface{ StackTrace() []byte }                  (samber/oops)
//     2. interface{ StackTrace() pkgerrors.StackTrace }    (github.com/pkg/errors)
//
// Typed-nil errors (a non-nil interface holding a nil pointer, e.g.
// (*MyErr)(nil) cast to error) are handled safely: error.type is emitted,
// error.message = "<nil>", and Error() is never called on the nil receiver.
//
// Returns zap.Skip() if err is nil so the caller can pass the result
// unconditionally:
//
//	logger.Error("operation failed",
//	    ServiceName("auth-api"),
//	    EventAction("user.login"),
//	    Err(err),
//	)
//
// The output JSON is flat dotted ECS keys (error.message, error.type,
// error.stack_trace) — identical to manually composing single-field helpers,
// but as one Field so it composes naturally with sibling fields.
func Err(err error) zapcore.Field {
	if err == nil {
		return zap.Skip()
	}
	return zap.Inline(errMarshaler{err: err})
}

// ErrAny returns a single inline zap.Field that emits ECS error.* fields from
// any value, intended for cases where the input is not statically typed as
// error — most commonly the result of recover() during panic handling.
// Behavior by input type:
//
//   - nil:             returns zap.Skip() (no fields emitted)
//   - typed-nil error: error.type emitted, error.message = "<nil>" — never
//     calls Error() on the typed-nil receiver, which would panic
//   - error:           same fields as Err(err), including error.stack_trace
//     when the error implements either StackTrace() []byte (samber/oops) or
//     StackTrace() pkgerrors.StackTrace (github.com/pkg/errors)
//   - other:           error.message = fmt.Sprint(v); error.type = fmt.Sprintf("%T", v)
//
// ErrAny intentionally does not call runtime/debug.Stack() itself. To attach
// the panic stack, pass it explicitly — callers may want to skip the cost or
// use a different stack source.
//
// Typical panic recovery:
//
//	defer func() {
//	    if r := recover(); r != nil {
//	        logger.Error("panic recovered",
//	            ErrAny(r),
//	            ErrorStackTrace(debug.Stack()),
//	        )
//	    }
//	}()
func ErrAny(v any) zapcore.Field {
	if v == nil {
		return zap.Skip()
	}
	return zap.Inline(errAnyMarshaler{v: v})
}

// errMarshaler renders an error into ECS error.* keys at the encoder's
// current namespace (no nesting). Used by Err via zap.Inline.
//
// Guards against typed-nil errors (non-nil interface holding a nil pointer)
// by checking before calling Error() — that call would otherwise panic on
// the nil receiver.
type errMarshaler struct{ err error }

func (m errMarshaler) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if isTypedNil(m.err) {
		enc.AddString("error.message", "<nil>")
		enc.AddString("error.type", fmt.Sprintf("%T", m.err))
		return nil
	}
	enc.AddString("error.message", m.err.Error())
	enc.AddString("error.type", fmt.Sprintf("%T", m.err))
	if stack := extractStackTrace(m.err); len(stack) > 0 {
		enc.AddByteString("error.stack_trace", stack)
	}
	return nil
}

// errAnyMarshaler renders any value into ECS error.* keys, routing error
// values through errMarshaler (which handles typed-nil) and falling back to
// fmt.Sprint / fmt.Sprintf("%T") for non-error values. Used by ErrAny via
// zap.Inline.
type errAnyMarshaler struct{ v any }

func (m errAnyMarshaler) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if err, ok := m.v.(error); ok {
		return errMarshaler{err: err}.MarshalLogObject(enc)
	}
	enc.AddString("error.message", fmt.Sprint(m.v))
	enc.AddString("error.type", fmt.Sprintf("%T", m.v))
	return nil
}

// extractStackTrace walks the error chain and returns the first stack trace
// found, in []byte form ready for ErrorStackTrace. Returns nil if no error in
// the chain carries a stack trace.
func extractStackTrace(err error) []byte {
	var bytesST stackTracerBytes
	if errors.As(err, &bytesST) {
		if s := bytesST.StackTrace(); len(s) > 0 {
			return s
		}
	}
	var pcsST stackTracerPCs
	if errors.As(err, &pcsST) {
		if s := pcsST.StackTrace(); len(s) > 0 {
			// pkg/errors.StackTrace implements fmt.Formatter; %+v renders each
			// frame as "function\n\tfile:line", matching what users expect to
			// see in error.stack_trace.
			return fmt.Appendf(nil, "%+v", s)
		}
	}
	return nil
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
