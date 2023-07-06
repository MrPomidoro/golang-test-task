package repository

import "github.com/golang-test-task/pkg/common/errors"

var ErrNotFound = errors.New("not found")
var ErrNothingInserted = errors.New("nothing inserted")
var ErrNothingDeleted = errors.New("nothing deleted")
var ErrNothingUpdated = errors.New("nothing updated")

//var ErrIntegrityConstraintViolation = errors.New("unique violation")
