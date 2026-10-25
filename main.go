package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("file open error: %v", err)
		return
	}
	defer f.Close()

	var currentLine string
	buf := make([]byte, 8)
	for {
		var chunk string
		n, err := f.Read(buf)
		if n > 0 {
			chunk = string(buf[:n])
		}
		parts := strings.Split(chunk, "\n")
		for i := 0; i < len(parts) - 1; i++ {
			fmt.Printf("read: %s\n", currentLine + parts[i])
			currentLine = ""
		}
		currentLine += parts[len(parts) - 1]
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("file read error: %v", err)
			return
		}

	}
	if currentLine != "" {
		fmt.Printf("read: %s\n", currentLine)
	}
}
