package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parserCat() (string, error) {
	args := os.Args[1:]
	var err error = nil

	if len(args) != 1 {
		err = fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
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
	filePath, err := parserCat()
	check(err)
	printFile(filePath)
}
