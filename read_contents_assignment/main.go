package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error when reading the file", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// f.Read(bs)
	// fmt.Println(string(bs))

	io.Copy(os.Stdout, f)
}
