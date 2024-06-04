package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func sha1Sum(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer r.Close()

	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}

func main() {
	sum, _ := sha1Sum("sha1/http.log.gz")
	fmt.Println(sum)
}
