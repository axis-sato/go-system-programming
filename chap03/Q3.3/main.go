package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Create("test.zip")
	if err != nil {
		print(err)
	}
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	writer, err := zipWriter.Create("newfile.txt")
	io.Copy(writer, strings.NewReader("foo"))
}
