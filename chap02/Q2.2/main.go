package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("test.csv")
	if err != nil {
		print(err)
	}
	writer := csv.NewWriter(file)
	writer.Write([]string{"a", "b"})
	writer.Write([]string{"c", "d"})
	writer.Flush()
}
