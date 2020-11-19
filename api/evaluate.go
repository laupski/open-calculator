package api

import (
	"encoding/json"
	"fmt"
	"github.com/laupski/open-calculator/internal/arithmetic"
	"net/http"
)

type evaluate struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
	Error  string `json:"error,omitempty"`
}

func (ch *calculatorHandlers) evaluate(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		query := r.URL.Query()
		input := query.Get("input")
		tokenList, err := arithmetic.Tokenize(input)
		output := arithmetic.ToTokenQueue(tokenList)
		sy := arithmetic.NewShuntingYard(output)
		postfixQueue := sy.InfixToPostFix()
		answer, err := arithmetic.Evaluate(postfixQueue)
		evaluate := evaluate{
			Input:  input,
			Output: fmt.Sprintf("%v", answer),
			Error:  fmt.Sprintf("%v", err),
		}

		jsonBytes, _ := json.Marshal(evaluate)
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
