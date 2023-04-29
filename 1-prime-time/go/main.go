package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net"
)

type Request struct {
	Method *string
	Number *float64
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

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
	reader := bufio.NewReader(conn)

	for {
		var request *Request

		bytes, err := reader.ReadBytes(byte('\n'))

		if err != nil {
			if err != io.EOF {
				fmt.Println("failed to read data, err:", err)
			}
			return
		}

		err = json.Unmarshal(bytes, &request)

		if err != nil {
			// this will handle invalid types
			fmt.Println("error unmarshalling json")
			conn.Write(bytes)
			return
		}

		if request.Method == nil || request.Number == nil {
			fmt.Println("method or number field is missing")
			// return malformed request and disconnect
			conn.Write(bytes)
			return
		}

		if *request.Method != "isPrime" {
			fmt.Println("invalid value for method")
			// return malformed request and disconnect
			conn.Write(bytes)
			return
		}

		// at this point, request is definitely valid and can be parsed
		fmt.Printf("request: %s %f\n", *request.Method, *request.Number)
		result := isPrime(*request.Number)

		response := &Response{
			Method: *request.Method,
			Prime:  result,
		}

		jsonBytes, err := json.Marshal(response)

		jsonBytes = append(jsonBytes, "\n"...)

		if err != nil {
			fmt.Println("error marshalling json")
			conn.Write(bytes)
			return
		}

		conn.Write(jsonBytes)
	}
}

func isPrime(n float64) bool {
	if n != float64(int(n)) {
		return false
	}

	if n < 2 {
		return false
	}

	sq_root := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq_root; i++ {
		if int(n)%i == 0 {
			fmt.Println("Non Prime Number")
			return false
		}
	}
	return true
}
