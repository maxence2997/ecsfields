// log_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestLog_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"LogLevel", ecszap.LogLevel("info"), "log.level", "info"},
		{"LogLogger", ecszap.LogLogger("kafka.consumer"), "log.logger", "kafka.consumer"},
		{"LogOriginFunction", ecszap.LogOriginFunction("pkg.Func"), "log.origin.function", "pkg.Func"},
		{"LogOriginFileName", ecszap.LogOriginFileName("foo.go"), "log.origin.file.name", "foo.go"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestLogOriginal(t *testing.T) {
	f := ecszap.LogOriginal([]byte("raw log line"))
	assert.Equal(t, "log.original", f.Key)
	assert.Equal(t, zapcore.ByteStringType, f.Type)
}

func TestLogOriginFileLine(t *testing.T) {
	f := ecszap.LogOriginFileLine(42)
	assert.Equal(t, "log.origin.file.line", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(42), f.Integer)
}
