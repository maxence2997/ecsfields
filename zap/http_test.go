// http_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestHTTP_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"HTTPRequestMethod", ecszap.HTTPRequestMethod("GET"), "http.request.method", "GET"},
		{"HTTPRequestID", ecszap.HTTPRequestID("req-1"), "http.request.id", "req-1"},
		{"HTTPRequestMimeType", ecszap.HTTPRequestMimeType("application/json"), "http.request.mime_type", "application/json"},
		{"HTTPRequestReferrer", ecszap.HTTPRequestReferrer("https://e.example/"), "http.request.referrer", "https://e.example/"},
		{"HTTPResponseMimeType", ecszap.HTTPResponseMimeType("text/html"), "http.response.mime_type", "text/html"},
		{"HTTPVersion", ecszap.HTTPVersion("1.1"), "http.version", "1.1"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestHTTP_BodyContent(t *testing.T) {
	r := ecszap.HTTPRequestBodyContent("hello")
	assert.Equal(t, "http.request.body.content", r.Key)
	assert.Equal(t, zapcore.StringType, r.Type)
	assert.Equal(t, "hello", r.String)

	s := ecszap.HTTPResponseBodyContent("world")
	assert.Equal(t, "http.response.body.content", s.Key)
	assert.Equal(t, "world", s.String)
}

func TestHTTP_LongFields(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantInt int64
	}{
		{"HTTPRequestBytes", ecszap.HTTPRequestBytes(1024), "http.request.bytes", 1024},
		{"HTTPRequestBodyBytes", ecszap.HTTPRequestBodyBytes(512), "http.request.body.bytes", 512},
		{"HTTPResponseBytes", ecszap.HTTPResponseBytes(2048), "http.response.bytes", 2048},
		{"HTTPResponseBodyBytes", ecszap.HTTPResponseBodyBytes(1024), "http.response.body.bytes", 1024},
		{"HTTPResponseStatusCode", ecszap.HTTPResponseStatusCode(200), "http.response.status_code", 200},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.Int64Type, tc.field.Type)
			assert.Equal(t, tc.wantInt, tc.field.Integer)
		})
	}
}
