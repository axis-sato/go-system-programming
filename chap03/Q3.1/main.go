package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("old.txt")
	if err != nil {
		print(err)
	}
	defer file.Close()
	newFile, err := os.Create("new.txt")
	if err != nil {
		print(err)
	}
	io.Copy(newFile, file)
}
