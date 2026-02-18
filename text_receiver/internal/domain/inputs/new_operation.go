package inputs

type NewOperation struct {
	Text string `json:"text" validate:"required" example:"test message 123 456 789 test"`
}
