package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func wcArgumentsValidator() (string, error) {

	args := flag.Args()
	var err error = nil

	if len(args) != 1 {
		err = fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
}

func count(fileName string) (int, int, int) {

	f, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(f)
	check(err)
	var linesCount int
	var wordCount int
	var charCount int
	for scanner.Scan() {
		linesCount++
		line := scanner.Text()
		charCount += len(line)
		wordCount += len(strings.Fields(line))
	}
	return linesCount, wordCount, charCount
}
func main() {
	var numberOfLinesOnly bool
	flag.BoolVar(&numberOfLinesOnly, "l", false, "Print only number of lines")
	var numberOfCharsOnly bool
	flag.BoolVar(&numberOfCharsOnly, "c", false, "Print only number of chars")
	var numberOfWordsOnly bool
	flag.BoolVar(&numberOfWordsOnly, "w", false, "Print only number of words")

	flag.Parse()
	filePath, err := wcArgumentsValidator()
	check(err)
	if !numberOfCharsOnly && !numberOfLinesOnly && !numberOfWordsOnly {
		numberOfCharsOnly = true
		numberOfLinesOnly = true
		numberOfWordsOnly = true
	}
	linesCount, wordCount, charCount := count(filePath)
	if numberOfLinesOnly {
		fmt.Print(linesCount)
	}
	if numberOfWordsOnly {
		fmt.Print(wordCount)
	}
	if numberOfCharsOnly {
		fmt.Print(charCount)
	}
}
