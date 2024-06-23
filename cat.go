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

func parserCat(input string) (string, error) {
	words := strings.Fields(input)
	if len(words) == 0 || words[0] != "cat" {
		return "", fmt.Errorf("error: command must start with 'cat'")
	}
	var filePath string = words[1]

	if len(words) != 2 {
		return "", fmt.Errorf("error: incorrect number of arguments")
	}

	return filePath, nil
}
func printFile(fileName string) {

	f, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(f)
	check(err)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, err := reader.ReadString('\n')
	check(err)
	filePath, err := parserCat(input)
	check(err)
	printFile(filePath)
}
