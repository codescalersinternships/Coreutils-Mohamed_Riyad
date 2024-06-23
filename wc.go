package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func parserWc(input string) (int, string, error) {
	words := strings.Fields(input)
	if len(words) == 0 || words[0] != "wc" {
		return 0, "", fmt.Errorf("error: command must start with 'wc'")
	}

	var num int = 0
	var filePath string = words[1]

	if len(words) == 3 {
		filePath = words[2]
		if words[1] == "-l" {
			num = 1
		} else if words[1] == "-w" {
			num = 2
		} else if words[1] == "-c" {
			num = 3
		} else {
			return 0, "", fmt.Errorf("error: incorrect flag arguments")
		}
	} else if len(words) != 2 {
		return 0, "", fmt.Errorf("error: incorrect number of arguments")
	}

	return num, filePath, nil
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
func printCount(linesCount int, wordCount int, charCount int, num int) {
	if num == 0 {
		fmt.Println("number of lines", linesCount)
		fmt.Println("number of words", wordCount)
		fmt.Println("number of chars", charCount)
	} else if num == 1 {
		fmt.Println("number of lines", linesCount)
	} else if num == 2 {
		fmt.Println("number of words", wordCount)
	} else if num == 3 {
		fmt.Println("number of chars", charCount)
	}
}
func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, err := reader.ReadString('\n')
	check(err)
	num, filePath, err := parserWc(input)
	check(err)
	linesCount, wordCount, charCount := count(filePath)
	printCount(linesCount, wordCount, charCount, num)
}
