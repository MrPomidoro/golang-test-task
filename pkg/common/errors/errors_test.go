package errors

import (
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	msg := "test error message"
	err := New(msg)

	assert.Equal(t, msg, err.Error(), "Error message should match")
}

func TestAs(t *testing.T) {
	err := Wrap(New("test error"), "wrapped")
	testError := New("test error")

	assert.True(t, As(testError, &err), "Should be able to cast error")
}

func TestIs(t *testing.T) {
	testError := New("test error")

	assert.True(t, Is(testError, testError), "Should be the same error")
}

func TestUnwrap(t *testing.T) {
	testError := New("test error")
	wrappedError := Wrap(testError, "wrapped")

	assert.NotNil(t, wrappedError, "Wrapped error should not be nil")
	assert.Contains(t, wrappedError.Error(), testError.Error(), "Wrapped error should contain the original error message")
}

func TestWrap(t *testing.T) {
	testError := New("test error")
	wrappedError := Wrap(testError, "wrapped")

	assert.NotNil(t, wrappedError, "Wrapped error should not be nil")
	assert.Contains(t, wrappedError.Error(), testError.Error(), "Wrapped error should contain the original error message")
}

func TestCause(t *testing.T) {
	testError := New("test error")
	wrappedError := Wrap(testError, "wrapped")

	assert.Equal(t, testError, Cause(wrappedError), "Should return the original error")
}

func TestAppend(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")

	err := Append(err1, err2)

	assert.Equal(t, 2, len(err.Errors), "Should combine the errors correctly")
}

func TestFlatten(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")
	err := Flatten(Append(err1, err2)).(*multierror.Error)

	assert.Equal(t, 2, len(err.Errors), "Should flatten the errors correctly")
}

func TestPrefix(t *testing.T) {
	err := New("error")
	prefix := "prefix:"
	prefixedErr := Prefix(err, prefix)

	assert.Equal(t, prefix+" "+err.Error(), prefixedErr.Error(), "Should add the prefix correctly")
}
