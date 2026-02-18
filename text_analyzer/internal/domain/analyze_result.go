package domain

type AnalyzeResult struct {
	WordCount         int     `json:"word_count" validate:"required"`
	SymbolsCount      int     `json:"symbols_count" validate:"required"`
	SentencesCount    int     `json:"sentences_count" validate:"required"`
	AverageWordLength float64 `json:"average_word_length" validate:"required"`
}
