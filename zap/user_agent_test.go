// user_agent_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestUserAgent(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"UserAgentOriginal", ecszap.UserAgentOriginal("Mozilla/5.0"), "user_agent.original", "Mozilla/5.0"},
		{"UserAgentName", ecszap.UserAgentName("Chrome"), "user_agent.name", "Chrome"},
		{"UserAgentVersion", ecszap.UserAgentVersion("120.0"), "user_agent.version", "120.0"},
		{"UserAgentDeviceName", ecszap.UserAgentDeviceName("iPhone"), "user_agent.device.name", "iPhone"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}
