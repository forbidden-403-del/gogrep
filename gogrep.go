package main

import (
	// "io"
	"os"
	"log"
	"fmt"
)

func getFileSize(filename string) int64 {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size()
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide a filename")
		return
	}
	filename := args[1]
	size := getFileSize(filename)
	data := make([]byte, size)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Read", count, "bytes")
	fmt.Println(string(data[:count]))
}