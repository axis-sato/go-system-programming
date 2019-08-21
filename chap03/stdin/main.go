package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err != nil {
			print(err)
		}
		fmt.Printf("size=%d input '%s'\n", size, string(buffer))
	}
}
