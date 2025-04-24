package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/fatih/color"
)

func getFileSize(filename string) int64 {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	return fi.Size()
}

func matchString(haystack, needle string) {
	if needle == "" {
		fmt.Println("No string to match")
		return
	}
	if len(needle) > len(haystack) {
		fmt.Println("String to match is longer than the haystack")
		return
	}

	index := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	for index != -1 {
		fmt.Print(color.YellowString(haystack[:index]))
		fmt.Print(color.GreenString(haystack[index:index+len(needle)]))
		haystack = haystack[index+len(needle):]
		index = strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	}
	fmt.Println(color.YellowString(haystack))
}

func main() {
	args := os.Args
	if len(args) < 3 {
		log.Fatal("./gogrep <filename> <string-to-match>")
		return
	}
	filename := args[1]
	to_match := args[2]

	size := getFileSize(filename)
	data := make([]byte, size)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	file.Read(data)
	string_data := string(data)
	matchString(string_data, to_match)
}
