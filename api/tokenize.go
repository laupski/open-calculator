package api

import (
	"encoding/json"
	"fmt"
	"github.com/laupski/open-calculator/internal/arithmetic"
	"net/http"
)

type tokenize struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

func (ch *calculatorHandlers) tokenize(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		input := query.Get("input")
		output, err := arithmetic.Tokenize(input)
		tokenize := tokenize{
			Input:  input,
			Output: fmt.Sprintf("%v", output),
			Error:  fmt.Sprintf("%v", err),
		}

		jsonBytes, _ := json.Marshal(tokenize)
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(jsonBytes)
		return
	default:
		tokenize := tokenize{Error: fmt.Sprintf("%v not allowed", r.Method)}
		jsonBytes, _ := json.Marshal(tokenize)
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = rw.Write(jsonBytes)
		return
	}
}
