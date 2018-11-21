package api

import (
	"fmt"
	"net/http"
)

// GetHome is the main endpoint
func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[INFO]", "Entering endpoint "+r.URL.RequestURI())

	// returns json with Response representation
	ResponseJSON(w, ToResponse("There's no place like home"))
}
