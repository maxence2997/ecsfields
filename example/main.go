// example/main.go — runnable demo of the recommended pattern: pair
// go.elastic.co/ecszap (envelope: level / ts / msg / caller / stacktrace)
// with ecsfields (typed user-added fields) for end-to-end ECS compliance.
//
// Want a minimal demo without ecszap? Replace the encoder/core/logger lines
// in main() with `logger := zap.Must(zap.NewProduction())`. ecsfields still
// works on its own — only the envelope keys (level/ts/msg) will fall back to
// zap's defaults instead of ECS-compliant names.

package main

import (
	"errors"
	"os"
	"time"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"

	ecsf "github.com/maxence2997/ecsfields/zap"
)

func main() {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.InfoLevel)
	logger := zap.New(core, zap.AddCaller())
	defer func() { _ = logger.Sync() }()

	start := time.Now().Add(-150 * time.Millisecond)
	err := errors.New("token expired")

	fields := []zap.Field{
		ecsf.ServiceName("auth-api"),
		ecsf.ServiceVersion("1.4.2"),
		ecsf.ServiceEnvironment("production"),

		ecsf.HostName("auth-api-7c9d"),

		ecsf.EventKind(ecsf.EventKindEvent),
		ecsf.EventCategory(ecsf.EventCategoryAuthentication),
		ecsf.EventType(ecsf.EventTypeStart, ecsf.EventTypeEnd),
		ecsf.EventOutcome(ecsf.EventOutcomeFailure),
		ecsf.EventAction("user.login"),
		ecsf.EventDuration(time.Since(start)),

		ecsf.HTTPRequestMethod("POST"),
		ecsf.HTTPResponseStatusCode(401),
		ecsf.URLPath("/api/v1/login"),
		ecsf.URLDomain("auth.example.com"),
		ecsf.ClientIP("203.0.113.7"),
		ecsf.UserAgentOriginal("Mozilla/5.0"),

		ecsf.TraceID("0af7651916cd43dd8448eb211c80319c"),
		ecsf.SpanID("b7ad6b7169203331"),

		ecsf.Label("tenant", "acme"),
		ecsf.Tags("login", "audit"),
	}
	fields = append(fields, ecsf.Err(err)...)

	logger.Info("login attempt failed", fields...)
}
