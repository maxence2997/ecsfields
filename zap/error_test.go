// error_test.go

package zap_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

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

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, "goroutine 1", enc.Fields["error.stack_trace"])
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

func TestErr_Nil(t *testing.T) {
	got := ecszap.Err(nil)
	assert.Nil(t, got)
}

func TestErr_StdlibError_NoStackTrace(t *testing.T) {
	err := errors.New("boom")
	got := ecszap.Err(err)
	require.Len(t, got, 2)

	keys := []string{got[0].Key, got[1].Key}
	assert.ElementsMatch(t, []string{"error.message", "error.type"}, keys)

	for _, f := range got {
		switch f.Key {
		case "error.message":
			assert.Equal(t, "boom", f.String)
		case "error.type":
			assert.Equal(t, fmt.Sprintf("%T", err), f.String)
		}
	}
}

type fakeStackTracer struct {
	msg   string
	stack []byte
}

func (f *fakeStackTracer) Error() string      { return f.msg }
func (f *fakeStackTracer) StackTrace() []byte { return f.stack }

func TestErr_StackTracer_EmitsStackTrace(t *testing.T) {
	err := &fakeStackTracer{msg: "kaboom", stack: []byte("goroutine 1...")}
	got := ecszap.Err(err)
	require.Len(t, got, 3)

	var keys []string
	for _, f := range got {
		keys = append(keys, f.Key)
	}
	assert.ElementsMatch(t,
		[]string{"error.message", "error.type", "error.stack_trace"},
		keys,
	)

	for _, f := range got {
		switch f.Key {
		case "error.message":
			assert.Equal(t, "kaboom", f.String)
		case "error.stack_trace":
			assert.Equal(t, zapcore.ByteStringType, f.Type)
			enc := zapcore.NewMapObjectEncoder()
			f.AddTo(enc)
			assert.Equal(t, "goroutine 1...", enc.Fields["error.stack_trace"])
		}
	}
}

func TestErrAny_Nil(t *testing.T) {
	assert.Nil(t, ecszap.ErrAny(nil))
}

func TestErrAny_Error_DelegatesToErr(t *testing.T) {
	err := errors.New("boom")
	got := ecszap.ErrAny(err)
	require.Len(t, got, 2)

	keys := []string{got[0].Key, got[1].Key}
	assert.ElementsMatch(t, []string{"error.message", "error.type"}, keys)
	for _, f := range got {
		switch f.Key {
		case "error.message":
			assert.Equal(t, "boom", f.String)
		case "error.type":
			assert.Equal(t, fmt.Sprintf("%T", err), f.String)
		}
	}
}

func TestErrAny_StackTracerError_IncludesStackTrace(t *testing.T) {
	err := &fakeStackTracer{msg: "kaboom", stack: []byte("goroutine 1...")}
	got := ecszap.ErrAny(err)
	require.Len(t, got, 3)

	var keys []string
	for _, f := range got {
		keys = append(keys, f.Key)
	}
	assert.ElementsMatch(t,
		[]string{"error.message", "error.type", "error.stack_trace"},
		keys,
	)
}

func TestErrAny_String(t *testing.T) {
	got := ecszap.ErrAny("oops")
	require.Len(t, got, 2)

	for _, f := range got {
		switch f.Key {
		case "error.message":
			assert.Equal(t, "oops", f.String)
		case "error.type":
			assert.Equal(t, "string", f.String)
		default:
			t.Fatalf("unexpected key %q", f.Key)
		}
	}
}

func TestErrAny_Int(t *testing.T) {
	got := ecszap.ErrAny(42)
	require.Len(t, got, 2)

	for _, f := range got {
		switch f.Key {
		case "error.message":
			assert.Equal(t, "42", f.String)
		case "error.type":
			assert.Equal(t, "int", f.String)
		default:
			t.Fatalf("unexpected key %q", f.Key)
		}
	}
}

type panicPayload struct {
	Reason string
}

func TestErrAny_Struct(t *testing.T) {
	got := ecszap.ErrAny(panicPayload{Reason: "deadlocked"})
	require.Len(t, got, 2)

	for _, f := range got {
		switch f.Key {
		case "error.message":
			assert.Equal(t, "{deadlocked}", f.String)
		case "error.type":
			assert.Equal(t, "zap_test.panicPayload", f.String)
		default:
			t.Fatalf("unexpected key %q", f.Key)
		}
	}
}

// derefingErr panics on Error() if the receiver is nil — emulating the common
// Go gotcha where panic(typedNilError) is recovered as a non-nil error
// interface holding a nil pointer.
type derefingErr struct{ msg string }

func (d *derefingErr) Error() string { return d.msg }

func TestErrAny_TypedNilPointerError_DoesNotPanic(t *testing.T) {
	var typedNil *derefingErr
	var asInterface error = typedNil

	var got []zapcore.Field
	require.NotPanics(t, func() {
		got = ecszap.ErrAny(asInterface)
	})
	require.Len(t, got, 2)

	for _, f := range got {
		switch f.Key {
		case "error.message":
			assert.Equal(t, "<nil>", f.String)
		case "error.type":
			assert.Equal(t, "*zap_test.derefingErr", f.String)
		default:
			t.Fatalf("unexpected key %q", f.Key)
		}
	}
}
