package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	fmt.Println("Listening port 8080: ")
	for {
		con, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go receiveFromConnection(con)
	}
}

func receiveFromConnection(con net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := con.Read(buffer)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(buffer[:n]))
	}
}
