package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	port := "8080"

	PrintMessage(port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

// PrintMessage shows the address the server is running
func PrintMessage(port string) {
	fmt.Println("Listening...")
	fmt.Println("URL: http://localhost:" + port)
	fmt.Println("------------------------------")
}
