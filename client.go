package main

import (
	"fmt"
	"net"
)

func main() {
	con, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	for {
		source := make([]byte, 1024)
		fmt.Print("Input to send: ")
		fmt.Scan(&source)
		con.Write(source)
	}
}
