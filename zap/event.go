// event.go — ECS event.* fields and typed enums for kind / outcome /
// category / type.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-event.html
//
// Out of scope (security domain, per spec):
//   event.risk_score, event.risk_score_norm, event.agent_id_status

package zap

import (
	"strconv"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// EventKindValue is one of the ECS-defined values for event.kind.
type EventKindValue string

// ECS-defined values for event.kind (ECS 8.17).
const (
	EventKindAlert         EventKindValue = "alert"
	EventKindAsset         EventKindValue = "asset"
	EventKindEnrichment    EventKindValue = "enrichment"
	EventKindEvent         EventKindValue = "event"
	EventKindMetric        EventKindValue = "metric"
	EventKindPipelineError EventKindValue = "pipeline_error"
	EventKindSignal        EventKindValue = "signal"
	EventKindState         EventKindValue = "state"
)

// EventOutcomeValue is one of the ECS-defined values for event.outcome.
type EventOutcomeValue string

// ECS-defined values for event.outcome (ECS 8.17).
const (
	EventOutcomeFailure EventOutcomeValue = "failure"
	EventOutcomeSuccess EventOutcomeValue = "success"
	EventOutcomeUnknown EventOutcomeValue = "unknown"
)

// EventCategoryValue is one of the ECS-recommended values for event.category.
// event.category is a keyword array; pass any combination to EventCategory.
type EventCategoryValue string

// ECS-recommended values for event.category (ECS 8.17).
const (
	EventCategoryAuthentication     EventCategoryValue = "authentication"
	EventCategoryConfiguration      EventCategoryValue = "configuration"
	EventCategoryDatabase           EventCategoryValue = "database"
	EventCategoryDriver             EventCategoryValue = "driver"
	EventCategoryEmail              EventCategoryValue = "email"
	EventCategoryFile               EventCategoryValue = "file"
	EventCategoryHost               EventCategoryValue = "host"
	EventCategoryIAM                EventCategoryValue = "iam"
	EventCategoryIntrusionDetection EventCategoryValue = "intrusion_detection"
	EventCategoryMalware            EventCategoryValue = "malware"
	EventCategoryNetwork            EventCategoryValue = "network"
	EventCategoryPackage            EventCategoryValue = "package"
	EventCategoryProcess            EventCategoryValue = "process"
	EventCategoryRegistry           EventCategoryValue = "registry"
	EventCategorySession            EventCategoryValue = "session"
	EventCategoryThreat             EventCategoryValue = "threat"
	EventCategoryVulnerability      EventCategoryValue = "vulnerability"
	EventCategoryWeb                EventCategoryValue = "web"
)

// EventTypeValue is one of the ECS-recommended values for event.type.
// event.type is a keyword array; pass any combination to EventType.
type EventTypeValue string

// ECS-recommended values for event.type (ECS 8.17).
const (
	EventTypeAccess       EventTypeValue = "access"
	EventTypeAdmin        EventTypeValue = "admin"
	EventTypeAllowed      EventTypeValue = "allowed"
	EventTypeChange       EventTypeValue = "change"
	EventTypeConnection   EventTypeValue = "connection"
	EventTypeCreation     EventTypeValue = "creation"
	EventTypeDeletion     EventTypeValue = "deletion"
	EventTypeDenied       EventTypeValue = "denied"
	EventTypeEnd          EventTypeValue = "end"
	EventTypeError        EventTypeValue = "error"
	EventTypeGroup        EventTypeValue = "group"
	EventTypeIndicator    EventTypeValue = "indicator"
	EventTypeInfo         EventTypeValue = "info"
	EventTypeInstallation EventTypeValue = "installation"
	EventTypeProtocol     EventTypeValue = "protocol"
	EventTypeStart        EventTypeValue = "start"
	EventTypeUser         EventTypeValue = "user"
)

// EventAction emits ECS event.action.
func EventAction(v string) zapcore.Field { return zap.String("event.action", v) }

// EventCategory emits ECS event.category as a keyword array.
func EventCategory(values ...EventCategoryValue) zapcore.Field {
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = string(v)
	}
	return zap.Strings("event.category", out)
}

// EventCode emits ECS event.code (keyword).
func EventCode(v string) zapcore.Field { return zap.String("event.code", v) }

// EventCodeInt is a convenience that formats a numeric code as an
// ECS event.code keyword.
func EventCodeInt(code int) zapcore.Field {
	return zap.String("event.code", strconv.Itoa(code))
}

// EventCreated emits ECS event.created.
func EventCreated(t time.Time) zapcore.Field { return zap.Time("event.created", t) }

// EventDataset emits ECS event.dataset.
func EventDataset(v string) zapcore.Field { return zap.String("event.dataset", v) }

// EventDuration emits ECS event.duration as nanoseconds (per ECS spec).
func EventDuration(d time.Duration) zapcore.Field {
	return zap.Int64("event.duration", d.Nanoseconds())
}

// EventEnd emits ECS event.end.
func EventEnd(t time.Time) zapcore.Field { return zap.Time("event.end", t) }

// EventHash emits ECS event.hash.
func EventHash(v string) zapcore.Field { return zap.String("event.hash", v) }

// EventID emits ECS event.id.
func EventID(v string) zapcore.Field { return zap.String("event.id", v) }

// EventIngested emits ECS event.ingested.
func EventIngested(t time.Time) zapcore.Field { return zap.Time("event.ingested", t) }

// EventKind emits ECS event.kind from a typed value.
func EventKind(v EventKindValue) zapcore.Field { return zap.String("event.kind", string(v)) }

// EventModule emits ECS event.module.
func EventModule(v string) zapcore.Field { return zap.String("event.module", v) }

// EventOriginal emits ECS event.original (raw bytes of the source event).
func EventOriginal(b []byte) zapcore.Field { return zap.ByteString("event.original", b) }

// EventOutcome emits ECS event.outcome from a typed value.
func EventOutcome(v EventOutcomeValue) zapcore.Field {
	return zap.String("event.outcome", string(v))
}

// EventProvider emits ECS event.provider.
func EventProvider(v string) zapcore.Field { return zap.String("event.provider", v) }

// EventReason emits ECS event.reason.
func EventReason(v string) zapcore.Field { return zap.String("event.reason", v) }

// EventReference emits ECS event.reference (URL pointing to additional info).
func EventReference(v string) zapcore.Field { return zap.String("event.reference", v) }

// EventSequence emits ECS event.sequence (long).
func EventSequence(seq int64) zapcore.Field { return zap.Int64("event.sequence", seq) }

// EventSeverity emits ECS event.severity (long).
func EventSeverity(s int64) zapcore.Field { return zap.Int64("event.severity", s) }

// EventStart emits ECS event.start.
func EventStart(t time.Time) zapcore.Field { return zap.Time("event.start", t) }

// EventType emits ECS event.type as a keyword array.
func EventType(values ...EventTypeValue) zapcore.Field {
	out := make([]string, len(values))
	for i, v := range values {
		out[i] = string(v)
	}
	return zap.Strings("event.type", out)
}
