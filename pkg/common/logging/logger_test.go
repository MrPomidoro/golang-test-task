package logging

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
	"testing"
	"time"
)

func TestL(t *testing.T) {
	logger := L(context.Background())
	assert.NotNil(t, logger, "Logger should not be nil")
}

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	assert.NotNil(t, logger, "Logger should not be nil")
}

func TestWithField(t *testing.T) {
	field := StringField("key", "value")
	obs, logs := observer.New(zap.InfoLevel)
	defaultLogger = zap.New(obs)

	WithField(context.Background(), field).Info("test")
	// Check that an Info-level Log was emitted with the correct message and fields.
	assert.Equal(t, 1, logs.Len(), "No log entry written")
	entry := logs.All()[0]
	assert.Equal(t, zap.InfoLevel, entry.Level, "Incorrect log level")
	assert.Contains(t, entry.Context, field, "Logger context should contain the field")
}

func TestWithFields(t *testing.T) {
	fields := []zap.Field{
		zap.String("key1", "value1"),
		zap.String("key2", "value2"),
	}

	obs, logs := observer.New(zap.InfoLevel)
	defaultLogger = zap.New(obs)

	WithFields(context.Background(), fields...).Info("test")

	// Check that an Info-level Log was emitted with the correct message and fields.
	assert.Equal(t, 1, logs.Len(), "No log entry written")
	entry := logs.All()[0]
	assert.Equal(t, zap.InfoLevel, entry.Level, "Incorrect log level")
	for _, field := range fields {
		assert.Contains(t, entry.Context, field, "Logger context should contain the field")
	}
}

func TestSetLevel(t *testing.T) {
	err := SetLevel("debug")
	assert.NoError(t, err, "Setting level to debug should not error")

	err = SetLevel("invalid")
	assert.Error(t, err, "Setting level to invalid should error")

	currentLevel := atomicLevel.Level()
	assert.Equal(t, zapcore.DebugLevel, currentLevel, "Current level should be DebugLevel")
}

func TestWithError(t *testing.T) {
	testErr := errors.New("test error")

	obs, logs := observer.New(zap.InfoLevel)
	defaultLogger = zap.New(obs)

	WithError(context.Background(), testErr).Info("test")

	// Check that an Info-level Log was emitted with the correct message and fields.
	assert.Equal(t, 1, logs.Len(), "No log entry written")
	entry := logs.All()[0]
	assert.Equal(t, zap.InfoLevel, entry.Level, "Incorrect log level")
	assert.Contains(t, entry.Context, zap.NamedError("error", testErr), "Logger context should contain the error")
}

func TestStringField(t *testing.T) {
	key := "testKey"
	value := "testValue"

	zField := zap.String(key, value)

	field := StringField(key, value)

	assert.Equal(t, key, field.Key, "Field key does not match the expected value")
	assert.Equal(t, value, field.String, "Field value does not match the expected value")
	assert.Equal(t, zField.Type, field.Type, "Field type is not zap.StringType")
}

func TestBoolField(t *testing.T) {
	key := "testKey"

	zField := zap.Bool(key, true)
	field := BoolField(key, true)

	assert.Equal(t, key, field.Key, "Field key does not match the expected value")
	assert.Equal(t, zField.Type, field.Type, "Field type is not zap.BoolType")

}

func TestDurationField(t *testing.T) {
	key := "testKey"
	value := time.Second

	zField := zap.Duration(key, value)
	field := DurationField(key, value)

	assert.Equal(t, key, field.Key, "Field key does not match the expected value")
	assert.Equal(t, zField.Type, field.Type, "Field type is not zap.DurationType")
}

func TestInt64Field(t *testing.T) {
	key := "testKey"
	value := int64(123)

	zField := zap.Int64(key, value)
	field := Int64Field(key, value)

	assert.Equal(t, key, field.Key, "Field key does not match the expected value")
	assert.Equal(t, zField.Integer, field.Integer, "Field type is not zap.Int64Type")
	assert.Equal(t, zField.Type, field.Type, "Field type is not zap.Int64Type")

}

func TestIntField(t *testing.T) {
	key := "testKey"
	value := 123

	zField := zap.Int(key, value)
	field := IntField(key, value)

	assert.Equal(t, key, field.Key, "Field key does not match the expected value")
	assert.Equal(t, zField.Integer, field.Integer, "Field type is not zap.Int64Type")
	assert.Equal(t, zField.Type, field.Type, "Field type is not zap.Int64Type")
}

func TestStringsField(t *testing.T) {
	key := "testKey"
	value := []string{"test1", "test2"}

	zField := zap.Strings(key, value)
	field := StringsField(key, value)

	assert.Equal(t, key, field.Key, "Field key does not match the expected value")
	assert.Equal(t, zField.String, field.String, "Field type is not zap.StringType")
	assert.Equal(t, zField.Type, field.Type, "Field type is not zap.StringType")
}
