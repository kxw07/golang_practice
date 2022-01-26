package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Needs to provide filename.")
		return
	}

	src, _ := os.Open(args[0])
	defer src.Close()

	// fileInfo, _ := src.Stat()
	// bs := make([]byte, fileInfo.Size())
	// src.Read(bs)
	// fmt.Println(string(bs))

	dst := os.Stdout
	defer dst.Close()

	n, _ := io.Copy(dst, src)
	fmt.Println("\nread bytes:", n)
}
