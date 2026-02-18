package errs

import "errors"

var ErrInvalidStatus = errors.New("invalid status value")

var ErrUnknownStauts = errors.New("unknown status")

var ErrOperationNotFound = errors.New("operation not found")
