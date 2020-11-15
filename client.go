package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

const (
	ConnType = "tcp"
	ConnPort = ":9999"
)

func main() {
	con, err := net.Dial(ConnType, ConnPort)

	if err != nil {
		fmt.Println("ERROR_CONNECTING: ", err)
	}

	defer con.Close()

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(con)

	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
			case nil:
				clientRequest := strings.TrimSpace(clientRequest)
				if _, err = con.Write([]byte(clientRequest + "\n")); err != nil {
					fmt.Printf("Failed to send the client request: %v\n", err)
				}
			case io.EOF:
				fmt.Println("Client closed the connection")
				return
			default:
				fmt.Printf("ERROR_CLIENT: %v\n", err)
				return
		}

		// Waiting for the server response
		serverResponse, err := serverReader.ReadString('\n')

		switch err {
			case nil:
				fmt.Println(strings.TrimSpace(serverResponse))
			case io.EOF:
				fmt.Println("Server closed the connection")
				return
			default:
				fmt.Printf("ERROR_SERVER: %v\n", err)
				return
		}
	}
}