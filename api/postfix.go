package api

import (
	"encoding/json"
	"fmt"
	"github.com/laupski/open-calculator/internal/arithmetic"
	"net/http"
)

type postfix struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

func (ch *calculatorHandlers) postfix(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		input := query.Get("input")
		tokenList, err := arithmetic.Tokenize(input)
		output := arithmetic.ToTokenQueue(tokenList)
		sy := arithmetic.NewShuntingYard(output)
		postfixQueue := sy.InfixToPostFix()
		postfix := tokenize{
			Input:  input,
			Output: fmt.Sprintf("%v", postfixQueue),
			Error:  fmt.Sprintf("%v", err),
		}

		jsonBytes, _ := json.Marshal(postfix)
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
