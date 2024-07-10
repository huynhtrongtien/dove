package log

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	instance Factory
)

func InitLogger(config *Configuration) error {
	if config.LogLevel == "" {
		config.LogLevel = DebugLevel
	}

	if config.StacktraceLevel == "" {
		config.StacktraceLevel = PanicLevel
	}

	if (config.File == nil) && (config.Console == nil) {
		return fmt.Errorf("log writer is nil")
	}

	instance = NewFactory(config)
	return nil
}

// Inst ...
func Inst() Factory {
	return instance
}

// Bg creates a context-unaware logger.
func Bg() Logger {
	return instance.Bg()
}

func WithContext(ctx context.Context) Logger {
	return instance.For(ctx)
}

// For returns a context-aware Logger. If the context
// contains an OpenTracing span, all logging calls are also
// echo-ed into the span.
func For(c *gin.Context) Logger {
	return instance.For(c.Request.Context())
}

func Field(key string, value interface{}) zapcore.Field {
	return zap.Any(key, value)
}

func Err(err error) zapcore.Field {
	return zap.Error(err)
}

// Debug logs an debig msg with fields
func Debug(c *gin.Context, msg string, fields ...zapcore.Field) {
	if c == nil {
		instance.Bg().Debug(msg, fields...)
		return
	}

	instance.For(c.Request.Context()).Debug(msg, fields...)
}

// Info logs an info msg with fields
func Info(c *gin.Context, msg string, fields ...zapcore.Field) {
	if c == nil {
		instance.Bg().Info(msg, fields...)
		return
	}

	instance.For(c.Request.Context()).Info(msg, fields...)
}

// Warn logs an warn msg with fields
func Warn(c *gin.Context, msg string, fields ...zapcore.Field) {
	if c == nil {
		instance.Bg().Warn(msg, fields...)
		return
	}

	instance.For(c.Request.Context()).Warn(msg, fields...)
}

// Error logs an error msg with fields
func Error(c *gin.Context, msg string, fields ...zapcore.Field) {
	if c == nil {
		instance.Bg().Error(msg, fields...)
		return
	}

	instance.For(c.Request.Context()).Error(msg, fields...)
}

// Fatal logs a fatal error msg with fields
func Fatal(c *gin.Context, msg string, fields ...zapcore.Field) {
	if c == nil {
		instance.Bg().Fatal(msg, fields...)
		return
	}

	instance.For(c.Request.Context()).Fatal(msg, fields...)
}

// Panic logs an panic msg with fields
func Panic(c *gin.Context, msg string, fields ...zapcore.Field) {
	if c == nil {
		instance.Bg().Panic(msg, fields...)
		return
	}

	instance.For(c.Request.Context()).Panic(msg, fields...)
}
