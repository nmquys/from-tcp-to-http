package main

import(
	"fmt"
	"os"
	"log"
	"bytes"
	"io"
)

func getLinesChannel(f io.ReadCloser) <- chan string {
	out := make(chan string, 1)

	go func(){
		defer f.Close()
		defer close(out)

		currentLine := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				currentLine += string(data[:i])
				data = data[i + 1:]
				out <- currentLine
				currentLine = ""
			}
			currentLine += string(data)
		}

		if len(currentLine) != 0 {
			out <- currentLine
		}
	}()

	return out
}

func main(){
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
