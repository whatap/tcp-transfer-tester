package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for i := 1; ; i += 1 {
		// Read message length
		var length uint32
		err := binary.Read(conn, binary.LittleEndian, &length)
		if err != nil {
			fmt.Println("Error reading message length:", err.Error())
			return
		}
		fmt.Printf("index: %d Received %d bytes header from client\n", i, length)

		// Read message
		buf := make([]byte, length)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			return
		}

		// Log received data
		fmt.Printf("\tRead %d bytes from client: %s\n", n, buf)

		err = binary.Write(conn, binary.LittleEndian, length)
		if err != nil {
			fmt.Println("Error writing message length:", err.Error())
			return
		}

		// Echo message back
		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}

		// Log sent data
		fmt.Printf("\tSent %d bytes to client: %s\n", n, buf)
	}
}

func main() {
	port := os.Getenv("SESSION_PORT")
	if len(port) < 1 {
		port = "6600"
	}
	ln, err := net.Listen("tcp", fmt.Sprint(":", port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on :", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		fmt.Println("Accepted connection from", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
