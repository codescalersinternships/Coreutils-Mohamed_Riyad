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
func parserWc() (string, error) {

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
	var linesCount int = 0
	var wordCount int = 0
	var charCount int = 0
	for scanner.Scan() {
		linesCount++
		line := scanner.Text()
		charCount += len(line)
		wordCount += len(strings.Fields(line))
	}
	return linesCount - 1, wordCount, charCount - 1
}
func main() {
	var numberOfLinesOnly bool
	flag.BoolVar(&numberOfLinesOnly, "l", false, "Print only number of lines")
	var numberOfCharsOnly bool
	flag.BoolVar(&numberOfCharsOnly, "c", false, "Print only number of chars")
	var numberOfWordsOnly bool
	flag.BoolVar(&numberOfWordsOnly, "w", false, "Print only number of words")

	flag.Parse()
	filePath, err := parserWc()
	check(err)
	linesCount, wordCount, charCount := count(filePath)
	if numberOfLinesOnly {
		fmt.Println("number of lines", linesCount)
	} else if numberOfWordsOnly {
		fmt.Println("number of words", wordCount)
	} else if numberOfCharsOnly {
		fmt.Println("number of chars", charCount)
	} else {
		fmt.Println("number of lines", linesCount)
		fmt.Println("number of words", wordCount)
		fmt.Println("number of chars", charCount)
	}
}
