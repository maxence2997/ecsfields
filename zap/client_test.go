// client_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestClient_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"ClientAddress", ecszap.ClientAddress("10.0.0.1"), "client.address", "10.0.0.1"},
		{"ClientIP", ecszap.ClientIP("10.0.0.1"), "client.ip", "10.0.0.1"},
		{"ClientMAC", ecszap.ClientMAC("aa:bb:cc:dd:ee:ff"), "client.mac", "aa:bb:cc:dd:ee:ff"},
		{"ClientDomain", ecszap.ClientDomain("e.example"), "client.domain", "e.example"},
		{"ClientRegisteredDomain", ecszap.ClientRegisteredDomain("example.com"), "client.registered_domain", "example.com"},
		{"ClientTopLevelDomain", ecszap.ClientTopLevelDomain("com"), "client.top_level_domain", "com"},
		{"ClientSubdomain", ecszap.ClientSubdomain("east"), "client.subdomain", "east"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestClientPort(t *testing.T) {
	f := ecszap.ClientPort(443)
	assert.Equal(t, "client.port", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(443), f.Integer)
}
