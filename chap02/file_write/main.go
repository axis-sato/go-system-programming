package main

import "os"

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		print(err)
	}
	file.Write([]byte("os.File exmple\n"))
	file.Close()
}
