package zap_test

import (
	"math"
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
	assert.Equal(t, 7.0, math.Float64frombits(uint64(f.Integer)))
}

func TestTags_Single(t *testing.T) {
	f := ecszap.Tags("auth", "critical")
	assert.Equal(t, "tags", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
	assert.NotNil(t, f.Interface)

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, []interface{}{"auth", "critical"}, enc.Fields["tags"])
}

func TestTags_Empty(t *testing.T) {
	f := ecszap.Tags()
	assert.Equal(t, "tags", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Empty(t, enc.Fields["tags"])
}
