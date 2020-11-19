package api

import (
	"fmt"
	"net/http"
)

const DefaultPort = 8080

// TODO add handler, routes and custom port
func StartApi(port int) {
	portstring := ":8080"
	fmt.Printf("Running open-calculator API on port: %v\n", port)
	err := http.ListenAndServe(portstring, nil)
	if err != nil {
		panic(err)
	}
}
