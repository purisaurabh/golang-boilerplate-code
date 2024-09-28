package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func SuccessResponse(rw http.ResponseWriter, status int, response interface{}) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling in success response:", err)
		status = http.StatusInternalServerError
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func ErrorResponse(rw http.ResponseWriter, status int, response interface{}) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error marshalling in error response:", err)
		status = http.StatusInternalServerError
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
