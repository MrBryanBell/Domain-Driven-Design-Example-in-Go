package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// SSE handler that sends periodic updates to connected clients
func sseHandler(w http.ResponseWriter, r *http.Request) {
	// Set SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Cache-Control")

	// Get the flusher to send data immediately
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Create a ticker to send events every 2 seconds
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	// Channel to detect client disconnect
	done := r.Context().Done()

	counter := 0
	for {
		select {
		case <-done:
			// Client disconnected
			log.Println("Client disconnected")
			return
		case <-ticker.C:
			counter++

			// Send different types of events
			if counter%3 == 0 {
				// Named event with JSON data
				fmt.Fprintf(w, "event: status\n")
				fmt.Fprintf(w, "data: {\"message\": \"System healthy\", \"count\": %d}\n\n", counter)
			} else {
				// Default event with simple message
				fmt.Fprintf(w, "data: Message %d sent at %s\n\n", counter, time.Now().Format("15:04:05"))
			}

			// Send the data immediately
			flusher.Flush()
		}
	}
}

// Serve a simple HTML page to test the SSE connection
func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Server-Sent Events Test</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        #messages { border: 1px solid #ccc; padding: 20px; height: 400px; overflow-y: auto; }
        .message { margin: 5px 0; padding: 5px; background: #f0f0f0; }
        .status { background: #e7f3ff; color: #0066cc; }
        button { padding: 10px 20px; margin: 10px 0; font-size: 16px; }
    </style>
</head>
<body>
    <h1>Server-Sent Events Demo</h1>
    <button onclick="startSSE()">Start Connection</button>
    <button onclick="stopSSE()">Stop Connection</button>
    <div id="status">Disconnected</div>
    <div id="messages"></div>

    <script>
        let eventSource = null;

        function startSSE() {
            if (eventSource) {
                eventSource.close();
            }

            eventSource = new EventSource('/events');
            
            eventSource.onopen = function(e) {
                document.getElementById('status').textContent = 'Connected';
                addMessage('Connection opened', 'status');
            };

            eventSource.onmessage = function(e) {
                addMessage('Default: ' + e.data, 'message');
            };

            eventSource.addEventListener('status', function(e) {
                addMessage('Status: ' + e.data, 'status');
            });

            eventSource.onerror = function(e) {
                document.getElementById('status').textContent = 'Error occurred';
                addMessage('Error occurred', 'message');
            };
        }

        function stopSSE() {
            if (eventSource) {
                eventSource.close();
                eventSource = null;
                document.getElementById('status').textContent = 'Disconnected';
                addMessage('Connection closed', 'status');
            }
        }

        function addMessage(message, className) {
            const messagesDiv = document.getElementById('messages');
            const messageDiv = document.createElement('div');
            messageDiv.className = 'message ' + className;
            messageDiv.textContent = new Date().toLocaleTimeString() + ': ' + message;
            messagesDiv.appendChild(messageDiv);
            messagesDiv.scrollTop = messagesDiv.scrollHeight;
        }
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func devHandler(w http.ResponseWriter, r *http.Request) {
	var html = `<!DOCTYPE html>
<html lang="es">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
	</head>
	<body>
		<h1>Hi there from Golang</h1>
	</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func main() {
	// Route handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/dev", devHandler)
	http.HandleFunc("/events", sseHandler)

	// Start server
	port := ":8080"
	log.Printf("Server starting on http://localhost%s", port)
	log.Printf("Open your browser and go to http://localhost:8080")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
