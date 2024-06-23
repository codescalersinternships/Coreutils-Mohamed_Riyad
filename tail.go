package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parserTail(input string) (int, string, error) {
	words := strings.Fields(input)
	if len(words) == 0 || words[0] != "tail" {
		return 0, "", fmt.Errorf("error: command must start with 'tail'")
	}

	var num int = 10
	var filePath string
	var err error

	if len(words) == 3 {
		words[1] = words[1][1:]
		num, err = strconv.Atoi(words[1])
		if err != nil {
			return 0, "", fmt.Errorf("error: second argument must be a number")
		}
		filePath = words[2]
	} else if len(words) == 2 {
		filePath = words[1]
	} else {
		return 0, "", fmt.Errorf("error: incorrect number of arguments")
	}

	return num, filePath, nil
}

func printLinesTail(fileName string, num int) {

	f, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(f)
	check(err)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = lines[0 : len(lines)-1]
	var n int = len(lines)

	for i := n - 1; i >= 0 && n-i <= num; i-- {
		fmt.Println(lines[i])
	}

}

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, err := reader.ReadString('\n')
	check(err)
	num, filePath, err := parserTail(input)
	check(err)
	printLinesTail(filePath, num)

}
