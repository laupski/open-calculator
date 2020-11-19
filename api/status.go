package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type status struct {
	Status string `json:"status,omitempty"`
	Error string `json:"error,omitempty"`
}

func (ch *calculatorHandlers) status(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		status := status{ Status: "API is up and running"}
		jsonBytes, _ := json.Marshal(status)
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(jsonBytes)
		return
	default:
		status := status{ Error: fmt.Sprintf("%v not allowed", r.Method) }
		jsonBytes, _ := json.Marshal(status)
		rw.Header().Add("content-type", "application/json")
		rw.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = rw.Write(jsonBytes)
		return
	}
}