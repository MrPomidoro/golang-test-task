package repository

import "github.com/golang-test-task/pkg/common/errors"

var ErrNotFound = errors.New("not found")
var ErrNothingInserted = errors.New("nothing inserted")
var ErrIntegrityConstraintViolation = errors.New("unique violation")
