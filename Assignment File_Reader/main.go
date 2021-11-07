package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No filename to read specified. Usage: $ go run main.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading from", filename)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	io.Copy(os.Stdout, fh)
}
