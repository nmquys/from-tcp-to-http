package main

import(
	"fmt"
	"os"
	"log"
)

func main(){
	f, err := os.Open("message.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		data := make([]byte, 8)
		readByte, err := f.Read(data)

		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(data[:readByte]))
	}
}
