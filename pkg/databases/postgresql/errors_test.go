package postgresql

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrCommit(t *testing.T) {
	testErr := errors.New("test error")
	wrappedErr := ErrCommit(testErr)

	assert.Contains(t, wrappedErr.Error(), "failed to commit Tx", "Error message does not contain expected substring")
	assert.True(t, errors.Is(wrappedErr, testErr), "Original error is not preserved")
}

func TestErrCreateQuery(t *testing.T) {
	testErr := errors.New("test error")
	wrappedErr := ErrCreateQuery(testErr)

	assert.Contains(t, wrappedErr.Error(), "failed to create SQL Query", "Error message does not contain expected substring")
	assert.True(t, errors.Is(wrappedErr, testErr), "Original error is not preserved")
}

func TestErrCreateTx(t *testing.T) {
	testErr := errors.New("test error")
	wrappedErr := ErrCreateTx(testErr)

	assert.Contains(t, wrappedErr.Error(), "failed to create Tx", "Error message does not contain expected substring")
	assert.True(t, errors.Is(wrappedErr, testErr), "Original error is not preserved")
}

func TestErrDoQuery(t *testing.T) {
	testErr := errors.New("test error")
	wrappedErr := ErrDoQuery(testErr)

	assert.Contains(t, wrappedErr.Error(), "failed to query", "Error message does not contain expected substring")
	assert.True(t, errors.Is(wrappedErr, testErr), "Original error is not preserved")
}

func TestErrExec(t *testing.T) {
	testErr := errors.New("test error")
	wrappedErr := ErrExec(testErr)

	assert.Contains(t, wrappedErr.Error(), "failed to execute", "Error message does not contain expected substring")
	assert.True(t, errors.Is(wrappedErr, testErr), "Original error is not preserved")
}

func TestErrRollback(t *testing.T) {
	testErr := errors.New("test error")
	wrappedErr := ErrRollback(testErr)

	assert.Contains(t, wrappedErr.Error(), "failed to rollback Tx", "Error message does not contain expected substring")
	assert.True(t, errors.Is(wrappedErr, testErr), "Original error is not preserved")
}

func TestErrScan(t *testing.T) {
	testErr := errors.New("test error")
	wrappedErr := ErrScan(testErr)

	assert.Contains(t, wrappedErr.Error(), "failed to scan", "Error message does not contain expected substring")
	assert.True(t, errors.Is(wrappedErr, testErr), "Original error is not preserved")
}
