package domain

type Status string

const (
	StatusInProcess Status = "in_process"
	StatusCompleted Status = "completed"
)

type Result struct {
	WordCount         int     `json:"word_count" validate:"required"`
	SymbolsCount      int     `json:"symbols_count" validate:"required"`
	SentencesCount    int     `json:"sentences_Count" validate:"required"`
	AverageWordLength float64 `json:"average_word_length" validate:"required"`
}

type Operation struct {
	ID     string  `json:"id" validate:"required"`
	Status Status  `json:"status" validate:"required"`
	Result *Result `json:"result"`
}

type ToProcessOperation struct {
	ID   string
	Text string
}

type SaveResponse struct {
	ID string `json:"id" validate:"required"`
}
