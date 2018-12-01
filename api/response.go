package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/eliamartani/go.blog/model"
)

/*
 * Public methods
 */

// HasError stops the process if any error is found and returns a message to the user
func HasError(err error) bool {
	if err != nil {
		fmt.Println("[WARNING] ", err.Error())

		return true
	}

	return false
}

// NoDataFoundResponse return a JSON object when no data is found
func NoDataFoundResponse(w http.ResponseWriter, message string) {
	if message == "" {
		message = "No data found"
	}

	responseJSON(w, model.Response{Code: http.StatusOK, Message: message})
}

// OKResponse return JSON representation from Response with status OK
func OKResponse(w http.ResponseWriter, message string) {
	if message == "" {
		message = "OK"
	}

	responseJSON(w, model.Response{Code: http.StatusOK, Message: message})
}

// OKDataResponse return JSON representation from Response with status OK
func OKDataResponse(w http.ResponseWriter, message string, data interface{}) {
	if message == "" {
		message = "OK"
	}

	responseJSON(w, model.Response{Code: http.StatusOK, Message: message, Data: data})
}

// PanicIfError stops the process if any error is found
func PanicIfError(err error) {
	if err != nil {
		// panic on the streets of london...
		panic(err.Error())
	}
}

// ServerErrorResponse return a JSON object when something bad happens
func ServerErrorResponse(w http.ResponseWriter, message string) {
	if message == "" {
		message = "Failed to load data"
	}

	responseJSON(w, model.Response{Code: http.StatusInternalServerError, Message: message})
}

/*
 * Private methods
 */

// enableCors allows external access for the api
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Length,Content-Type")
}

// responseJSON tries to handle messages to the client
func responseJSON(w http.ResponseWriter, response model.Response) {
	w.Header().Set("Content-Type", "application/json")

	enableCors(&w)

	w.WriteHeader(response.Code)

	json.NewEncoder(w).Encode(response)
}
