package main

import (
	"io"
	"strings"
	"fmt"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		defer f.Close()

		var currentLine string
		buf := make([]byte, 8)
		for {
			var chunk string
			n, err := f.Read(buf)
			if n > 0 {
				chunk = string(buf[:n])
				parts := strings.Split(chunk, "\n")
				for i := 0; i < len(parts) - 1; i++ {
					out <- currentLine + parts[i]
					currentLine = ""
				}
				currentLine += parts[len(parts) - 1]
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("file read error: %v\n", err)
				break
			}
		}
		if currentLine != "" {
			out <- currentLine
		}
	}()

	return out
}
