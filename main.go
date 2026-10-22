package main

import (
	"io"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("file open error: %v", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 8)
	for {
		b, err := f.Read(buf)
		if b > 0 {
			fmt.Printf("read: %s\n", buf[:b])
		}
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("file read error: %v", err)
			return
		}

	}
}
