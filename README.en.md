# TCP Server and Client Implemented in Go

## Overview

This repository contains two Go programs implementing simple TCP server and client communication. The server echoes back any messages it receives from clients, while the client sends messages to the server and displays the server's response.

## TCP Server

### Description

The TCP server program listens for incoming connections on a specified port. Upon client connection, it reads messages from the client, echoes them back, and logs the received and sent data. The server supports concurrent connections by handling each client in a separate goroutine.

### Usage

To run the TCP server, execute the server.go file. By default, the server listens on port 6600, but you can customize the port by setting the SESSION_PORT environment variable.

```bash
export SESSION_PORT=6600
$ ./server 
Server listening on : 6600
Accepted connection from 127.0.0.1:41190
index: 1 Received 18 bytes header from client
Read 18 bytes from client: hello world whatap
Sent 18 bytes to client: hello world whatap
```


## TCP Client

### Description

The TCP client program connects to a specified server IP address and port. It prompts the user to enter a message to send to the server. After sending the message, it waits for the server's response and displays it. This process is logged along with timing information.

### Usage

To run the TCP client, execute the client.go file. By default, the client connects to 127.0.0.1:6600, but you can customize the server IP address and port by setting the SESSION_SERVER_IP and SESSION_SERVER_PORT environment variables, respectively.

```bash
export SESSION_SERVER_IP="127.0.0.1"
export SESSION_SERVER_PORT="6600"
$ ./client
Enter message(default: hello world whatap):
2024-05-29 11:32:06.039883098 +0900 KST m=+2.425403694 Connected to server
index:1  Sent 18 bytes to server in 22.433µs
Received 18 bytes header from server in 132.983µs
Read 18 bytes from server in 140.527µs: hello world whatap
connection time:  270.856µs
index:2  Sent 18 bytes to server in 69.973µs
Received 18 bytes header from server in 113.095µs
Read 18 bytes from server in 127.112µs: hello world whatap
connection time:  1.000676571s
```

