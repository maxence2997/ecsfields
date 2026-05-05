// trace_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestTrace_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"TraceID", ecszap.TraceID("abc"), "trace.id", "abc"},
		{"SpanID", ecszap.SpanID("def"), "span.id", "def"},
		{"TransactionID", ecszap.TransactionID("ghi"), "transaction.id", "ghi"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}
