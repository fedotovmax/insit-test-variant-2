package usecases

import (
	"context"
	"log/slog"
	"unicode"

	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/domain"
	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/domain/inputs"
)

type AnalyzeUsecase struct {
	log *slog.Logger
}

func NewAnalyzeUsecase(log *slog.Logger,
) *AnalyzeUsecase {
	return &AnalyzeUsecase{
		log: log,
	}
}

func (u *AnalyzeUsecase) Execute(ctx context.Context, text inputs.Text) *domain.AnalyzeResult {

	const op = "usecases.analyze"

	var (
		wordCount     int
		charCount     int
		sentenceCount int
		totalWordLen  int
		currentWord   int
		inSentenceEnd bool
	)

	runes := []rune(text.Data)

	n := len(runes)

	for i := 0; i < n; i++ {
		r := runes[i]
		charCount++

		// проверка на конец по многоточие
		if i+2 < n && runes[i] == '.' && runes[i+1] == '.' && runes[i+2] == '.' {
			if !inSentenceEnd {
				sentenceCount++
				inSentenceEnd = true
			}
			charCount += 2
			i += 2
			continue
		}

		// проверка на конц предложния на ? или !
		if r == '!' || r == '?' {
			if !inSentenceEnd {
				sentenceCount++
				inSentenceEnd = true
			}
		} else if r == '.' {
			// проверка сокращений (т.д. т.п.)
			isAbbr := false
			if i > 0 && i+1 < n {
				prev := runes[i-1]
				next := runes[i+1]
				if unicode.IsLetter(prev) && unicode.IsLower(next) {
					isAbbr = true
				}
			}
			if !inSentenceEnd && !isAbbr {
				sentenceCount++
				inSentenceEnd = true
			}
		} else {
			inSentenceEnd = false
		}

		// считаем слова
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			currentWord++
		} else {
			if currentWord > 0 {
				wordCount++
				totalWordLen += currentWord
				currentWord = 0
			}
		}
	}

	// tсли текст закончился словом
	if currentWord > 0 {
		wordCount++
		totalWordLen += currentWord
		currentWord = 0
	}

	var avg float64
	if wordCount > 0 {
		avg = float64(totalWordLen) / float64(wordCount)
	}

	result := &domain.AnalyzeResult{
		WordCount:         wordCount,
		SymbolsCount:      charCount,
		SentencesCount:    sentenceCount,
		AverageWordLength: avg,
	}

	return result
}
