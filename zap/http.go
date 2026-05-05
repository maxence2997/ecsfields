// http.go — ECS http.* fields.
//
// ECS reference (8.17): https://www.elastic.co/guide/en/ecs/8.17/ecs-http.html

package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HTTPRequestMethod emits ECS http.request.method.
func HTTPRequestMethod(v string) zapcore.Field { return zap.String("http.request.method", v) }

// HTTPRequestBodyBytes emits ECS http.request.body.bytes.
func HTTPRequestBodyBytes(n int64) zapcore.Field {
	return zap.Int64("http.request.body.bytes", n)
}

// HTTPRequestBodyContent emits ECS http.request.body.content.
func HTTPRequestBodyContent(v string) zapcore.Field {
	return zap.String("http.request.body.content", v)
}

// HTTPRequestBytes emits ECS http.request.bytes.
func HTTPRequestBytes(n int64) zapcore.Field { return zap.Int64("http.request.bytes", n) }

// HTTPRequestID emits ECS http.request.id.
func HTTPRequestID(v string) zapcore.Field { return zap.String("http.request.id", v) }

// HTTPRequestMimeType emits ECS http.request.mime_type.
func HTTPRequestMimeType(v string) zapcore.Field {
	return zap.String("http.request.mime_type", v)
}

// HTTPRequestReferrer emits ECS http.request.referrer.
func HTTPRequestReferrer(v string) zapcore.Field {
	return zap.String("http.request.referrer", v)
}

// HTTPResponseBodyBytes emits ECS http.response.body.bytes.
func HTTPResponseBodyBytes(n int64) zapcore.Field {
	return zap.Int64("http.response.body.bytes", n)
}

// HTTPResponseBodyContent emits ECS http.response.body.content.
func HTTPResponseBodyContent(v string) zapcore.Field {
	return zap.String("http.response.body.content", v)
}

// HTTPResponseBytes emits ECS http.response.bytes.
func HTTPResponseBytes(n int64) zapcore.Field { return zap.Int64("http.response.bytes", n) }

// HTTPResponseMimeType emits ECS http.response.mime_type.
func HTTPResponseMimeType(v string) zapcore.Field {
	return zap.String("http.response.mime_type", v)
}

// HTTPResponseStatusCode emits ECS http.response.status_code.
func HTTPResponseStatusCode(code int) zapcore.Field {
	return zap.Int64("http.response.status_code", int64(code))
}

// HTTPVersion emits ECS http.version.
func HTTPVersion(v string) zapcore.Field { return zap.String("http.version", v) }
