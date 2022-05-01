package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error :[", err, "]")
		os.Exit(1)
	}

	//fmt.Println(resp)

	//a byte slice with 99999 elements
	//This is done because read function is not setup to increase the size of byte slice
	respByte := make([]byte, 99999)
	resp.Body.Read(respByte)
	//this will print the google html page
	fmt.Println(string(respByte))

	//o/p using writer interface
	//NOTE: Read() can be used only once on resp type
	// When line 24-27 are uncommented, io.Copy does not write anything to console.
	//TODO : read why this happens
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

//write bytes to console
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	fmt.Println("Bytes processed [", len(bs), "]")

	return len(bs), nil
}
