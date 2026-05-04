// server_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestServer_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"ServerAddress", ecszap.ServerAddress("10.0.0.2"), "server.address", "10.0.0.2"},
		{"ServerIP", ecszap.ServerIP("10.0.0.2"), "server.ip", "10.0.0.2"},
		{"ServerMAC", ecszap.ServerMAC("aa:bb:cc:dd:ee:ff"), "server.mac", "aa:bb:cc:dd:ee:ff"},
		{"ServerDomain", ecszap.ServerDomain("e.example"), "server.domain", "e.example"},
		{"ServerRegisteredDomain", ecszap.ServerRegisteredDomain("example.com"), "server.registered_domain", "example.com"},
		{"ServerTopLevelDomain", ecszap.ServerTopLevelDomain("com"), "server.top_level_domain", "com"},
		{"ServerSubdomain", ecszap.ServerSubdomain("east"), "server.subdomain", "east"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestServerPort(t *testing.T) {
	f := ecszap.ServerPort(8080)
	assert.Equal(t, "server.port", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(8080), f.Integer)
}
