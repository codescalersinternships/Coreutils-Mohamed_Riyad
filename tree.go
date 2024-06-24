package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func parserTree(input string) (string, error) {
	words := strings.Fields(input)
	if len(words) == 0 || words[0] != "tree" {
		return "", fmt.Errorf("error: command must start with 'tree'")
	}

	if len(words) != 2 {
		return "", fmt.Errorf("error: incorrect number of arguments")
	}

	var filePath string = words[1]

	return filePath, nil
}
func printSpaces(level int) {
	spaces := ""
	for i := 0; i < level; i++ {
		spaces += "  "
	}
	fmt.Print(spaces)
}
func printTree(dirPath string, level int) {
	fileInfo, err := os.Stat(dirPath)
	check(err)
	printSpaces(level)
	fmt.Println(fileInfo.Name())

	if !fileInfo.Mode().IsRegular() {
		dir, err := os.Open(dirPath)
		check(err)
		files, err := dir.Readdir(-1)
		check(err)
		for _, file := range files {
			newPath := filepath.Join(dirPath, file.Name())
			printTree(newPath, level+1)
		}
	}
}
func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, err := reader.ReadString('\n')
	check(err)
	dirPath, err := parserTree(input)
	check(err)
	printTree(dirPath, 0)
}
