package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		currentLine := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)

			if n > 0 {
				data = data[:n]
				dataStr := string(data)
				parts := strings.Split(dataStr, "\n")

				for i := 0; i < len(parts)-1; i++ {
					out <- currentLine + parts[i]
					currentLine = ""
				}

				currentLine += parts[len(parts)-1]
			}

			if err != nil {
				break
			}
		}

		if len(currentLine) != 0 {
			out <- currentLine
		}
	}()

	return out
}

const port = ":42069"

func main() {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error listening for TCP traffic: %s\n", err.Error())
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %s\n", err.Error())
		}

		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s\n", line)
		}
	}
}
