package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/random", getRandomNumber)
	http.HandleFunc("/", serveIndex)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func getRandomNumber(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s %s\n", r.Method, r.URL)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	// Generate random number
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100)

	// Create the response JSON
	response := fmt.Sprintf(`{"value": %d}`, randomNum)

	// Write the response
	w.Write([]byte(response))
}
