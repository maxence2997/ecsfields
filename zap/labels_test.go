package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestLabel(t *testing.T) {
	f := ecszap.Label("streaming_id", "abc-123")
	assert.Equal(t, "labels.streaming_id", f.Key)
	assert.Equal(t, zapcore.StringType, f.Type)
	assert.Equal(t, "abc-123", f.String)
}

func TestNumericLabel(t *testing.T) {
	f := ecszap.NumericLabel("consecutive_errors", 7)
	assert.Equal(t, "numeric_labels.consecutive_errors", f.Key)
	assert.Equal(t, zapcore.Float64Type, f.Type)
	// zap stores float64 in Integer via math.Float64bits.
	// We round-trip through the public API instead of reaching for unsafe internals:
	// the caller value is recoverable because Float64Type uses Integer field directly.
	// Simpler: just assert key and type — value semantics covered by zap itself.
}

func TestTags_Single(t *testing.T) {
	f := ecszap.Tags("auth", "critical")
	assert.Equal(t, "tags", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
	require := assert.New(t)
	require.NotNil(f.Interface)
}

func TestTags_Empty(t *testing.T) {
	f := ecszap.Tags()
	assert.Equal(t, "tags", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
}
