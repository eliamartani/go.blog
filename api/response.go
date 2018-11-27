package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response is a representation from object
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// PanicIfError stops the process if any error is found
func PanicIfError(err error) {
	if err != nil {
		// panic on the streets of london...
		panic(err.Error())
	}
}

// HasError stops the process if any error is found and returns a message to the user
func HasError(err error) bool {
	if err != nil {
		fmt.Println("[WARNING]", err.Error())

		return true
	}

	return false
}

// ToResponse creates an output handler
func ToResponse(jsonValue interface{}) Response {
	return Response{Code: http.StatusOK, Data: jsonValue}
}

// OK return a JSON object when the need is to return a message with Status OK
func OK(message string) Response {
	return Response{Code: http.StatusOK, Message: message}
}

// NoDataFound return a JSON object when no data is found
func NoDataFound() Response {
	return Response{Code: http.StatusOK, Message: "No data found"}
}

// ServerError return a JSON object when something bad happens
func ServerError() Response {
	return Response{Code: http.StatusInternalServerError, Message: "Failed to load data"}
}

// ResponseJSON tries to handle messages to the client
func ResponseJSON(w http.ResponseWriter, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)

	json.NewEncoder(w).Encode(response)
}
