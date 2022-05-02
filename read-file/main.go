package main

import (
	"fmt"
	"io"
	"os"
)

//TASK : Create a program that reads the contents of a text file
//then prints its contents to the terminal

func main() {
	filePath := os.Args[1]           //access file path
	file, error := os.Open(filePath) // open file on filepath
	if error != nil {                //check if there is error
		fmt.Println("Specified file not found. file-path=[", filePath, "] error=[", error, "]")
		os.Exit(1)
	}
	io.Copy(os.Stdout, file) // copy file.Read to console. File implements reader interface
}
