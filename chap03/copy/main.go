package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		print(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
