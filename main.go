package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("failed to accept req: %v", err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	buf, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalf("failed to read req: %v", err)
	}
	log.Printf("recived: %v", string(buf))

	_, err = conn.Write([]byte("received message\n"))
	if err != nil {
		log.Fatalf("failed to write response: %v", err)
	}
	return
}
