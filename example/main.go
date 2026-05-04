// example/main.go — runnable demo that emits an ECS-shaped log line using
// ecsfields and the standard zap JSON encoder.

package main

import (
	"errors"
	"time"

	"go.uber.org/zap"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer func() { _ = logger.Sync() }()

	start := time.Now().Add(-150 * time.Millisecond)
	err := errors.New("token expired")

	fields := []zap.Field{
		ecszap.ServiceName("auth-api"),
		ecszap.ServiceVersion("1.4.2"),
		ecszap.ServiceEnvironment("production"),

		ecszap.HostName("auth-api-7c9d"),

		ecszap.EventKind(ecszap.EventKindEvent),
		ecszap.EventCategory(ecszap.EventCategoryAuthentication),
		ecszap.EventType(ecszap.EventTypeStart, ecszap.EventTypeEnd),
		ecszap.EventOutcome(ecszap.EventOutcomeFailure),
		ecszap.EventAction("user.login"),
		ecszap.EventDuration(time.Since(start)),

		ecszap.HTTPRequestMethod("POST"),
		ecszap.HTTPResponseStatusCode(401),
		ecszap.URLPath("/api/v1/login"),
		ecszap.URLDomain("auth.example.com"),
		ecszap.ClientIP("203.0.113.7"),
		ecszap.UserAgentOriginal("Mozilla/5.0"),

		ecszap.TraceID("0af7651916cd43dd8448eb211c80319c"),
		ecszap.SpanID("b7ad6b7169203331"),

		ecszap.Label("tenant", "acme"),
		ecszap.Tags("login", "audit"),
	}
	fields = append(fields, ecszap.Err(err)...)

	logger.Info("login attempt failed", fields...)
}
