// error_test.go

package zap_test

import (
	"errors"
	"fmt"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

// emit applies the field to a MapObjectEncoder and returns the resulting key
// map. Inline fields (Err / ErrAny) flatten their sub-keys into this map.
func emit(f zapcore.Field) map[string]any {
	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	return enc.Fields
}

func TestErrorMessage(t *testing.T) {
	f := ecszap.ErrorMessage("boom")
	assert.Equal(t, "error.message", f.Key)
	assert.Equal(t, zapcore.StringType, f.Type)
	assert.Equal(t, "boom", f.String)
}

func TestErrorType(t *testing.T) {
	f := ecszap.ErrorType("*url.Error")
	assert.Equal(t, "error.type", f.Key)
	assert.Equal(t, zapcore.StringType, f.Type)
	assert.Equal(t, "*url.Error", f.String)
}

func TestErrorStackTrace(t *testing.T) {
	f := ecszap.ErrorStackTrace([]byte("goroutine 1"))
	assert.Equal(t, "error.stack_trace", f.Key)
	assert.Equal(t, zapcore.ByteStringType, f.Type)
	assert.Equal(t, "goroutine 1", emit(f)["error.stack_trace"])
}

func TestErrorCode(t *testing.T) {
	f := ecszap.ErrorCode("E_42")
	assert.Equal(t, "error.code", f.Key)
	assert.Equal(t, zapcore.StringType, f.Type)
	assert.Equal(t, "E_42", f.String)
}

func TestErrorID(t *testing.T) {
	f := ecszap.ErrorID("err-1")
	assert.Equal(t, "error.id", f.Key)
	assert.Equal(t, zapcore.StringType, f.Type)
	assert.Equal(t, "err-1", f.String)
}

func TestErr_Nil_ReturnsSkip(t *testing.T) {
	f := ecszap.Err(nil)
	assert.Equal(t, zapcore.SkipType, f.Type)
	assert.Empty(t, emit(f))
}

func TestErr_StdlibError_NoStackTrace(t *testing.T) {
	err := errors.New("boom")
	got := emit(ecszap.Err(err))

	assert.Equal(t, "boom", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", err), got["error.type"])
	assert.NotContains(t, got, "error.stack_trace")
}

type fakeStackTracer struct {
	msg   string
	stack []byte
}

func (f *fakeStackTracer) Error() string      { return f.msg }
func (f *fakeStackTracer) StackTrace() []byte { return f.stack }

func TestErr_BytesStackTracer_EmitsStackTrace(t *testing.T) {
	err := &fakeStackTracer{msg: "kaboom", stack: []byte("goroutine 1...")}
	got := emit(ecszap.Err(err))

	assert.Equal(t, "kaboom", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", err), got["error.type"])
	assert.Equal(t, "goroutine 1...", got["error.stack_trace"])
}

func TestErr_PkgErrorsStackTracer_EmitsStackTrace(t *testing.T) {
	err := pkgerrors.New("boom")
	got := emit(ecszap.Err(err))

	assert.Equal(t, "boom", got["error.message"])
	assert.NotEmpty(t, got["error.type"])

	stack, ok := got["error.stack_trace"].(string)
	require.True(t, ok, "error.stack_trace should be string-coercible")
	assert.NotEmpty(t, stack)
	assert.Contains(t, stack, "TestErr_PkgErrorsStackTracer_EmitsStackTrace",
		"stack should reference the test function frame")
}

func TestErrAny_Nil_ReturnsSkip(t *testing.T) {
	f := ecszap.ErrAny(nil)
	assert.Equal(t, zapcore.SkipType, f.Type)
	assert.Empty(t, emit(f))
}

func TestErrAny_Error_DelegatesToErr(t *testing.T) {
	err := errors.New("boom")
	got := emit(ecszap.ErrAny(err))

	assert.Equal(t, "boom", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", err), got["error.type"])
	assert.NotContains(t, got, "error.stack_trace")
}

func TestErrAny_StackTracerError_IncludesStackTrace(t *testing.T) {
	err := &fakeStackTracer{msg: "kaboom", stack: []byte("goroutine 1...")}
	got := emit(ecszap.ErrAny(err))

	assert.Equal(t, "kaboom", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", err), got["error.type"])
	assert.Equal(t, "goroutine 1...", got["error.stack_trace"])
}

func TestErrAny_String(t *testing.T) {
	v := "oops"
	got := emit(ecszap.ErrAny(v))

	assert.Equal(t, v, got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", v), got["error.type"])
}

func TestErrAny_Int(t *testing.T) {
	v := 42
	got := emit(ecszap.ErrAny(v))

	assert.Equal(t, "42", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", v), got["error.type"])
}

type panicPayload struct {
	Reason string
}

func TestErrAny_Struct(t *testing.T) {
	v := panicPayload{Reason: "deadlocked"}
	got := emit(ecszap.ErrAny(v))

	assert.Equal(t, "{deadlocked}", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", v), got["error.type"])
}

// derefingErr panics on Error() if the receiver is nil — emulating the common
// Go gotcha where panic(typedNilError) is recovered as a non-nil error
// interface holding a nil pointer.
type derefingErr struct{ msg string }

func (d *derefingErr) Error() string { return d.msg }

func TestErrAny_TypedNilPointerError_DoesNotPanic(t *testing.T) {
	var typedNil *derefingErr
	var asInterface error = typedNil

	var got map[string]any
	require.NotPanics(t, func() {
		got = emit(ecszap.ErrAny(asInterface))
	})

	assert.Equal(t, "<nil>", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", asInterface), got["error.type"])
}

func TestErr_TypedNilPointerError_DoesNotPanic(t *testing.T) {
	var typedNil *derefingErr
	var asInterface error = typedNil

	var got map[string]any
	require.NotPanics(t, func() {
		got = emit(ecszap.Err(asInterface))
	})

	assert.Equal(t, "<nil>", got["error.message"])
	assert.Equal(t, fmt.Sprintf("%T", asInterface), got["error.type"])
}
