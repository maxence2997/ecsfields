// url_test.go

package zap_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"

	ecszap "github.com/maxence2997/ecsfields/zap"
)

func TestURL_StringHelpers(t *testing.T) {
	cases := []struct {
		name    string
		field   zapcore.Field
		wantKey string
		wantStr string
	}{
		{"URLOriginal", ecszap.URLOriginal("https://e.example/p?q=1"), "url.original", "https://e.example/p?q=1"},
		{"URLFull", ecszap.URLFull("https://e.example/p"), "url.full", "https://e.example/p"},
		{"URLScheme", ecszap.URLScheme("https"), "url.scheme", "https"},
		{"URLDomain", ecszap.URLDomain("e.example"), "url.domain", "e.example"},
		{"URLRegisteredDomain", ecszap.URLRegisteredDomain("example.com"), "url.registered_domain", "example.com"},
		{"URLTopLevelDomain", ecszap.URLTopLevelDomain("com"), "url.top_level_domain", "com"},
		{"URLSubdomain", ecszap.URLSubdomain("east"), "url.subdomain", "east"},
		{"URLPath", ecszap.URLPath("/p"), "url.path", "/p"},
		{"URLQuery", ecszap.URLQuery("q=1"), "url.query", "q=1"},
		{"URLFragment", ecszap.URLFragment("hash"), "url.fragment", "hash"},
		{"URLUsername", ecszap.URLUsername("alice"), "url.username", "alice"},
		{"URLPassword", ecszap.URLPassword("s3cret"), "url.password", "s3cret"},
		{"URLExtension", ecszap.URLExtension("png"), "url.extension", "png"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantKey, tc.field.Key)
			assert.Equal(t, zapcore.StringType, tc.field.Type)
			assert.Equal(t, tc.wantStr, tc.field.String)
		})
	}
}

func TestURLPort(t *testing.T) {
	f := ecszap.URLPort(443)
	assert.Equal(t, "url.port", f.Key)
	assert.Equal(t, zapcore.Int64Type, f.Type)
	assert.Equal(t, int64(443), f.Integer)
}
