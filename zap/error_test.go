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
		}
	}
}
