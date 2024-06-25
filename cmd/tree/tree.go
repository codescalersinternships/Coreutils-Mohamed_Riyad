package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func parserTree() (string, error) {
	args := os.Args[1:]
	var err error = nil

	if len(args) != 1 {
		err = fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
}
func printSpaces(level int) {
	spaces := strings.Repeat("  ", level)
	fmt.Print(spaces + "|__")
}
func printTree(dirPath string, level int) {
	fileInfo, err := os.Stat(dirPath)
	check(err)
	if level != 0 {
		printSpaces(level)
	}
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
	dirPath, err := parserTree()
	check(err)
	printTree(dirPath, 0)
}
