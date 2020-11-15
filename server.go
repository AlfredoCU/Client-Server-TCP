package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	Type = "tcp"
	Port = ":9999"
)

func main() {
	listener, err := net.Listen(Type, Port)

	if err != nil {
		log.Fatalln("ERROR_LISTENING: ", err)
	}

	defer listener.Close()

	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR_CONNECTING: ", err)
			continue
		}

		// If you want, you can increment a counter here and inject to handleClientRequest below as client identifier
		go handleClientRequest(con)
	}
}

func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
			case nil:
				clientRequest := strings.TrimSpace(clientRequest)
				if clientRequest == ":QUIT" {
					fmt.Println("Client requested server to close the connection so closing")
					return
				} else {
					fmt.Println(clientRequest)
				}
			case io.EOF:
				fmt.Println("Client closed the connection by terminating the process")
				return
			default:
				fmt.Printf("ERROR: %v\n", err)
				return
		}

		// Responding to the client request
		if _, err = con.Write([]byte("GOT IT!\n")); err != nil {
			fmt.Printf("Dailed to respond to client: %v\n", err)
		}
	}
}