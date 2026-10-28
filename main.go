package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		fmt.Printf("file open error: %v", err)
		return
	}

	lines := getLinesChannel(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

}
