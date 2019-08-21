package main

import (
	"archive/zip"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=test.zip")

	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()
	writer, err := zipWriter.Create("foo.txt")
	if err != nil {
		print(err)
	}
	writer.Write([]byte("foo"))
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8888", nil)
}
