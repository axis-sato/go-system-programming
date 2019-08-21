package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var csvSource = `1, foo, ふー, フー
2, bar, ばー, バー`

func main() {
	reader := strings.NewReader(csvSource)
	csvReader := csv.NewReader(reader)
	for {
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(line[0], line[1:4])
	}
}
