package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type UserArguments struct {
	Filename string
	Flags    []string
}

type Count struct {
	wordCount int
	charCount int
}

func main() {
	wc := &Count{}
	var outputString string
	var uArgs UserArguments

	args := os.Args
	if len(args) < 3 {
		log.Fatal("please provide the file name! and flags")
	}

	for _, val := range args[1:] {
		if strings.Contains(val, "-") {
			uArgs.Flags = append(uArgs.Flags, val)
		} else {
			uArgs.Filename = val
		}
	}

	file, err := os.Open(uArgs.Filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		totalWords := strings.Fields(line)
		wc.wordCount += len(totalWords)
		wc.charCount += len(line)
	}

	for _, value := range uArgs.Flags {
		switch value {
		case "-c":
			outputString += fmt.Sprintf("Total character count: %d\n", wc.charCount)

		case "-w":
			outputString += fmt.Sprintf("Total word count: %d\n", wc.wordCount)
		}
	}

	fmt.Println(outputString)
}
