package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var wc int

	args := os.Args
	if len(args) < 2 {
		log.Fatal("please provide the file name!")
	}

	fileName := args[1:2]
	file, err := os.Open(fileName[0])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		totalWords := strings.Fields(line)
		wc += len(totalWords)
	}

	fmt.Printf("Total words in files: %d\n", wc)
}
