package main

import (
	"crypto/rand"
	"io"
	"os"
)

func copyN(dest io.Writer, src io.Reader, length int) {
	io.Copy(dest, io.LimitReader(src, int64(length)))
}

func main() {
	file, err := os.Create("test.bin")
	if err != nil {
		print(err)
	}
	defer file.Close()
	copyN(file, rand.Reader, 1024)
}
