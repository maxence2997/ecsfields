// process_test.go

package zap_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestProcess_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"ProcessCommandLine", ecszap.ProcessCommandLine("/bin/sh -c echo"), "process.command_line", "/bin/sh -c echo"},
		{"ProcessExecutable", ecszap.ProcessExecutable("/usr/bin/sh"), "process.executable", "/usr/bin/sh"},
		{"ProcessName", ecszap.ProcessName("sh"), "process.name", "sh"},
		{"ProcessTitle", ecszap.ProcessTitle("nginx: worker"), "process.title", "nginx: worker"},
		{"ProcessWorkingDirectory", ecszap.ProcessWorkingDirectory("/var/run"), "process.working_directory", "/var/run"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestProcessArgs(t *testing.T) {
	f := ecszap.ProcessArgs("/bin/sh", "-c", "echo hi")
	assert.Equal(t, "process.args", f.Key)
	assert.Equal(t, zapcore.ArrayMarshalerType, f.Type)
}

func TestProcessArgsCount(t *testing.T) {
	f := ecszap.ProcessArgsCount(3)
	assert.Equal(t, "process.args_count", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(3), f.Integer)
}

func TestProcessExitCode(t *testing.T) {
	f := ecszap.ProcessExitCode(137)
	assert.Equal(t, "process.exit_code", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(137), f.Integer)
}

func TestProcessPID(t *testing.T) {
	f := ecszap.ProcessPID(4321)
	assert.Equal(t, "process.pid", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(4321), f.Integer)
}

func TestProcessStart(t *testing.T) {
	now := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	f := ecszap.ProcessStart(now)
	assert.Equal(t, "process.start", f.Key)
	assert.Equal(t, zapcore.TimeType, f.Type)
}

func TestProcessUptime_Seconds(t *testing.T) {
	f := ecszap.ProcessUptime(2 * time.Minute)
	assert.Equal(t, "process.uptime", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(120), f.Integer)
}
