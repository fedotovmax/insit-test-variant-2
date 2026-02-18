package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/fedotovmax74/insit-test-variant-2/text_analyzer/internal/keys"
)

func DecodeJSON(r io.Reader, v any) error {
	defer io.Copy(io.Discard, r)
	return json.NewDecoder(r).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v any) {

	w.Header().Set(keys.HeaderContentType, keys.ContentTypeJSON)
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, `{"message": "failed to encode json"}`, http.StatusInternalServerError)
	}
}
