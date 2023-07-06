package logging

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

var (
	defaultLogger = NewLogger()
	atomicLevel   = zap.NewAtomicLevel()
)

func NewLogger() *zap.Logger {
	var encCfg zapcore.EncoderConfig

	var encoder zapcore.Encoder

	encCfg = zap.NewProductionEncoderConfig()
	encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encCfg.LevelKey = zapcore.DebugLevel.String()

	encCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encCfg.TimeKey = "timestamp"

	encoder = zapcore.NewConsoleEncoder(encCfg)

	l := zap.New(
		zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), atomicLevel),
	)
	l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return l
}

func GetLogger() *zap.Logger {
	return defaultLogger
}

func SetLevel(level string) error {
	var lvl zapcore.Level

	switch strings.ToLower(level) {
	case "debug":
		lvl = zapcore.DebugLevel
	case "info":
		lvl = zapcore.InfoLevel
	case "warn":
		lvl = zapcore.WarnLevel
	case "error":
		lvl = zapcore.ErrorLevel
	case "dpanic":
		lvl = zapcore.DPanicLevel
	case "panic":
		lvl = zapcore.PanicLevel
	case "fatal":
		lvl = zapcore.FatalLevel
	default:
		return fmt.Errorf("invalid log level: %s", level)
	}

	atomicLevel.SetLevel(lvl)
	return nil
}

func L(ctx context.Context) *zap.Logger {
	return LoggerFromContext(ctx)
}

func WithField(ctx context.Context, field zap.Field) *zap.Logger {
	return L(ctx).With(field)
}

func WithFields(ctx context.Context, fields ...zap.Field) *zap.Logger {
	return L(ctx).With(fields...)
}

func WithError(ctx context.Context, err error) *zap.Logger {
	return L(ctx).With(zap.Error(err))
}

func StringField(key, val string) zap.Field {
	return zap.String(key, val)
}

func IntField(key string, val int) zap.Field {
	return zap.Int(key, val)
}

func DurationField(key string, val time.Duration) zap.Field {
	return zap.Duration(key, val)
}

func Int64Field(key string, val int64) zap.Field {
	return zap.Int64(key, val)
}

func StringsField(key string, val []string) zap.Field {
	return zap.Strings(key, val)
}

func BoolField(key string, val bool) zap.Field {
	return zap.Bool(key, val)
}
