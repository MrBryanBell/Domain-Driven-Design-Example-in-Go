package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connection, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			log.Println("Error connecting with client...")
		}
		defer func() {
			log.Printf("Connection with client %s has been closed", connection.RemoteAddr().String())
			connection.Close()
		}()

		var ticker = time.NewTicker(2 * time.Second)
		defer ticker.Stop()

		var timeout = time.After(6 * time.Second)

		for {
			select {
			case <-ticker.C:
				connection.WriteMessage(websocket.TextMessage, []byte("We are connected "+time.Now().String()))
			case <-timeout:
				return
			}
		}
	})

	log.Println("Server is running in port 8080...")
	http.ListenAndServe(":8080", nil)
}
