package api

import (
	"fmt"
	"net/http"
	"strconv"
)

const DefaultPort = 8080

func StartApi(port int) {
	portString := ":" + strconv.Itoa(port)

	calculatorHandlers := newCalculatorHandlers()
	http.HandleFunc("/api/v1/tokenize", calculatorHandlers.tokenize)
	http.HandleFunc("/api/v1/postfix", calculatorHandlers.postfix)
	http.HandleFunc("/api/v1/evaluate", calculatorHandlers.evaluate)
	http.HandleFunc("/api/v1/status", calculatorHandlers.status)
	fmt.Printf("Running open-calculator API on port: %v\n", port)
	err := http.ListenAndServe(portString, nil)
	if err != nil {
		panic(err)
	}
}

type calculatorHandlers struct{}

func newCalculatorHandlers() *calculatorHandlers {
	return &calculatorHandlers{}
}
