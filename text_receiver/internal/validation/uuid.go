package validation

import (
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidUUIDFormat = errors.New("invalid uuid format")

func IsUUID(value string) error {

	_, err := uuid.Parse(value)

	if err != nil {
		return ErrInvalidUUIDFormat
	}

	return nil
}
