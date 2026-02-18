package inputs

import (
	"errors"
	"fmt"

	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/validation"
)

type Text struct {
	Data string `json:"data" validate:"required" example:"test message 123 456 789 test"`
}

func (i *Text) Validate() error {

	var validationErrors []error

	err := validation.EmptyString(i.Data)

	if err != nil {
		validationErrors = append(validationErrors, fmt.Errorf("%s: %w", "Data", err))
	}

	return errors.Join(validationErrors...)

}
