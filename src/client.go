package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	serverip := os.Getenv("SESSION_SERVER_IP")
	if len(serverip) < 1 {
		serverip = "127.0.0.1"
	}
	serverport := os.Getenv("SESSION_SERVER_PORT")
	if len(serverport) < 1 {
		serverport = "6600"
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter message(default: hello world whatap): ")
	scanner.Scan()
	message := scanner.Text()
	if len(message) < 1 {
		message = "hello world whatap"
	}

	t1Time := time.Now()
	conn, err := net.Dial("tcp", fmt.Sprint(serverip, ":", serverport))
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	connectedTime := time.Now()
	fmt.Println(connectedTime, "Connected to server in ", connectedTime.Sub(t1Time))

	for i := 1; ; i++ {

		// Record sending start time
		sendStartTime := time.Now()

		// Write message length
		length := uint32(len(message))
		err := binary.Write(conn, binary.LittleEndian, length)
		if err != nil {
			fmt.Println("Error writing message length:", err.Error())
			return
		}

		// Write message
		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing message:", err.Error())
			return
		}

		// Record sending end time
		sendEndTime := time.Now()

		// Calculate sending duration
		sendDuration := sendEndTime.Sub(sendStartTime)

		// Log sending data
		fmt.Printf("index:%d(%s)  Sent %d bytes to server in %s\n", i, time.Now().Sub(connectedTime), len(message), sendDuration)

		// Record response start time
		responseStartTime := time.Now()

		// Read server response length
		var responseLength uint32
		err = binary.Read(conn, binary.LittleEndian, &responseLength)
		if err != nil {
			fmt.Println("Error reading server response length:", err.Error())
			return
		}
		fmt.Printf("\tReceived %d bytes header from server in %s\n", responseLength, time.Now().Sub(responseStartTime))

		// Read server response
		responseBuf := make([]byte, responseLength)
		nbyteLeft := responseLength
		for nbyteLeft > 0 {
			n, err := conn.Read(responseBuf)
			if err != nil {
				fmt.Println("Error reading server response:", err.Error())
				return
			}
			nbyteLeft = nbyteLeft - uint32(n)
			// Record response end time
			responseEndTime := time.Now()

			// Calculate response duration
			responseDuration := responseEndTime.Sub(responseStartTime)

			// Log received data
			fmt.Printf("\tRead %d bytes from server in %s: %s\n", n, responseDuration, responseBuf)
		}

		time.Sleep(1 * time.Second)
	}
}
