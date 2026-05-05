// event_test.go

package zap_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestEvent_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"EventAction", ecszap.EventAction("user.login"), "event.action", "user.login"},
		{"EventCode", ecszap.EventCode("E_42"), "event.code", "E_42"},
		{"EventCodeInt", ecszap.EventCodeInt(7), "event.code", "7"},
		{"EventDataset", ecszap.EventDataset("auth.audit"), "event.dataset", "auth.audit"},
		{"EventHash", ecszap.EventHash("deadbeef"), "event.hash", "deadbeef"},
		{"EventID", ecszap.EventID("evt-1"), "event.id", "evt-1"},
		{"EventModule", ecszap.EventModule("auth"), "event.module", "auth"},
		{"EventProvider", ecszap.EventProvider("oauth2"), "event.provider", "oauth2"},
		{"EventReason", ecszap.EventReason("token_expired"), "event.reason", "token_expired"},
		{"EventReference", ecszap.EventReference("https://docs.example/E_42"), "event.reference", "https://docs.example/E_42"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestEventDuration_Nanoseconds(t *testing.T) {
	f := ecszap.EventDuration(time.Second)
	assert.Equal(t, "event.duration", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(1_000_000_000), f.Integer)
}

func TestEventOriginal_ByteString(t *testing.T) {
	payload := []byte(`{"x":1}`)
	f := ecszap.EventOriginal(payload)
	assert.Equal(t, "event.original", f.Key)
	assert.Equal(t, zapcore.ByteStringType, f.Type)

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, string(payload), enc.Fields["event.original"])
}

func TestEvent_TimeFields(t *testing.T) {
	now := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
	}{
		{"EventCreated", ecszap.EventCreated(now), "event.created"},
		{"EventEnd", ecszap.EventEnd(now), "event.end"},
		{"EventIngested", ecszap.EventIngested(now), "event.ingested"},
		{"EventStart", ecszap.EventStart(now), "event.start"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.TimeType, tc.field.Type)

			enc := zapcore.NewMapObjectEncoder()
			tc.field.AddTo(enc)
			assert.Equal(t, now, enc.Fields[tc.wantKey])
		})
	}
}

func TestEvent_LongFields(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantInt int64
	}{
		{"EventSequence", ecszap.EventSequence(42), "event.sequence", 42},
		{"EventSeverity", ecszap.EventSeverity(7), "event.severity", 7},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.Int64Type, tc.field.Type)
			assert.Equal(t, tc.wantInt, tc.field.Integer)
		})
	}
}

func TestEventKind_TypedEnum(t *testing.T) {
	f := ecszap.EventKind(ecszap.EventKindEvent)
	assert.Equal(t, "event.kind", f.Key)
	assert.Equal(t, zapcore.StringType, f.Type)
	assert.Equal(t, "event", f.String)

	// All ECS-defined event.kind values per ECS 8.17.
	want := map[ecszap.EventKindValue]string{
		ecszap.EventKindAlert:         "alert",
		ecszap.EventKindAsset:         "asset",
		ecszap.EventKindEnrichment:    "enrichment",
		ecszap.EventKindEvent:         "event",
		ecszap.EventKindMetric:        "metric",
		ecszap.EventKindPipelineError: "pipeline_error",
		ecszap.EventKindSignal:        "signal",
		ecszap.EventKindState:         "state",
	}
	for v, s := range want {
		assert.Equal(t, s, string(v))
	}
}

func TestEventOutcome_TypedEnum(t *testing.T) {
	f := ecszap.EventOutcome(ecszap.EventOutcomeSuccess)
	assert.Equal(t, "event.outcome", f.Key)
	assert.Equal(t, zapcore.StringType, f.Type)
	assert.Equal(t, "success", f.String)

	assert.Equal(t, "failure", string(ecszap.EventOutcomeFailure))
	assert.Equal(t, "success", string(ecszap.EventOutcomeSuccess))
	assert.Equal(t, "unknown", string(ecszap.EventOutcomeUnknown))
}

func TestEventCategory_TypedArray(t *testing.T) {
	f := ecszap.EventCategory(ecszap.EventCategoryAuthentication, ecszap.EventCategoryWeb)
	assert.Equal(t, "event.category", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
	assert.Equal(t, "authentication", string(ecszap.EventCategoryAuthentication))
	assert.Equal(t, "web", string(ecszap.EventCategoryWeb))

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, []interface{}{"authentication", "web"}, enc.Fields["event.category"])
}

func TestEventType_TypedArray(t *testing.T) {
	f := ecszap.EventType(ecszap.EventTypeStart, ecszap.EventTypeEnd)
	assert.Equal(t, "event.type", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
	assert.Equal(t, "start", string(ecszap.EventTypeStart))
	assert.Equal(t, "end", string(ecszap.EventTypeEnd))

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, []interface{}{"start", "end"}, enc.Fields["event.type"])
}
