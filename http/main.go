package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{} //this is a Writer Interface(check the docs(type Write interface{}))

func main() {
	resp, err := http.Get("https://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// bs := make([]byte, 99999)

	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
