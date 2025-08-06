package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	const PORT = ":8080"
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Error starting TCP Listener in port: %s", PORT)
	}
	log.Printf("Listening in port: %s \n", PORT)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("Successful connection with %s\n", conn.RemoteAddr().String())

	// write a greeting message when connection is established...
	// fmt.Fprintf(conn, "Hello connection from %s\n", conn.RemoteAddr().String())
	var greeting = fmt.Sprintf("Hello connection from %s\n", conn.RemoteAddr().String())
	conn.Write([]byte(greeting))

	// Why do we need to create a Buffer?
	// https://claude.ai/chat/e05786fe-8c47-4572-a64e-dbf40aa42226
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("Connection with %s closed: %v\n", conn.RemoteAddr().String(), err)
			return
		}
		message := string(buffer[:n])
		log.Printf("Received from %s: %s", conn.RemoteAddr().String(), message)

		// echo the message back to the client
		conn.Write([]byte("Received: " + message))
	}
}
