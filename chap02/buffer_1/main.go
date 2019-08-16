package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("byte buffer example\n"))
	buffer.Write([]byte("foo\n"))
	fmt.Println(buffer.String())
}
