package logging

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
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
	return defaultLogger
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
