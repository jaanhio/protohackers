package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":3333")

	if err != nil {
		fmt.Printf("error listening")
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.Fatal("error accepting connection")
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	io.Copy(conn, conn)
}
