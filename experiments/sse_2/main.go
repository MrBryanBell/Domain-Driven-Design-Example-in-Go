package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Greeting struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		// write a simple greeting message
		writer.Header().Set("Content-Type", "application/json")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.WriteHeader(http.StatusOK)

		greeting := Greeting{Message: "Hello, World!"}
		fmt.Fprintf(writer, `{"message": "%s"}`, greeting.Message)
	})

	http.HandleFunc("/see", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/event-stream")
		writer.Header().Set("Cache-Control", "no-cache")
		writer.Header().Set("Connection", "keep-alive")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		fmt.Println("New SSE connection established")

		flusher, ok := writer.(http.Flusher)
		if !ok {
			http.Error(writer, "Streaming is unsupported", http.StatusInternalServerError)
			return
		}

		for i := range 10 {
			// Stream Event should have a `data:` prefix
			fmt.Fprintf(writer, "data: count is %d\n\n", i)
			flusher.Flush()

			time.Sleep(1 * time.Second)
		}
	})

	log.Println("Listening on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting the server...")
	}
}
