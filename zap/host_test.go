// host_test.go

package zap_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestHost_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"HostArchitecture", ecszap.HostArchitecture("arm64"), "host.architecture", "arm64"},
		{"HostDomain", ecszap.HostDomain("example.com"), "host.domain", "example.com"},
		{"HostHostname", ecszap.HostHostname("worker-1.local"), "host.hostname", "worker-1.local"},
		{"HostID", ecszap.HostID("h-1"), "host.id", "h-1"},
		{"HostName", ecszap.HostName("worker-1"), "host.name", "worker-1"},
		{"HostType", ecszap.HostType("vm"), "host.type", "vm"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestHostIP(t *testing.T) {
	f := ecszap.HostIP("10.0.0.1", "fe80::1")
	assert.Equal(t, "host.ip", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
	assert.NotNil(t, f.Interface)

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, []interface{}{"10.0.0.1", "fe80::1"}, enc.Fields["host.ip"])
}

func TestHostMAC(t *testing.T) {
	f := ecszap.HostMAC("00-1B-44-11-3A-B7")
	assert.Equal(t, "host.mac", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)

	enc := zapcore.NewMapObjectEncoder()
	f.AddTo(enc)
	assert.Equal(t, []interface{}{"00-1B-44-11-3A-B7"}, enc.Fields["host.mac"])
}

func TestHostUptime_Seconds(t *testing.T) {
	f := ecszap.HostUptime(2 * time.Hour)
	assert.Equal(t, "host.uptime", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(7200), f.Integer)
}
