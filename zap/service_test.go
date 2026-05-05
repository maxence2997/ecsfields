// service_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestService_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"ServiceAddress", ecszap.ServiceAddress("10.0.0.1:8080"), "service.address", "10.0.0.1:8080"},
		{"ServiceEnvironment", ecszap.ServiceEnvironment("prod"), "service.environment", "prod"},
		{"ServiceEphemeralID", ecszap.ServiceEphemeralID("eph-1"), "service.ephemeral_id", "eph-1"},
		{"ServiceID", ecszap.ServiceID("svc-1"), "service.id", "svc-1"},
		{"ServiceName", ecszap.ServiceName("auth"), "service.name", "auth"},
		{"ServiceNodeName", ecszap.ServiceNodeName("node-a"), "service.node.name", "node-a"},
		{"ServiceNodeRole", ecszap.ServiceNodeRole("primary"), "service.node.role", "primary"},
		{"ServiceState", ecszap.ServiceState("running"), "service.state", "running"},
		{"ServiceType", ecszap.ServiceType("elasticsearch"), "service.type", "elasticsearch"},
		{"ServiceVersion", ecszap.ServiceVersion("1.2.3"), "service.version", "1.2.3"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestServiceNodeRoles(t *testing.T) {
	f := ecszap.ServiceNodeRoles("primary", "voting")
	assert.Equal(t, "service.node.roles", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
	assert.NotNil(t, f.Interface)

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, []interface{}{"primary", "voting"}, enc.Fields["service.node.roles"])
}
