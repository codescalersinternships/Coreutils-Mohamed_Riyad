package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func catArgumentsValidator() (string, error) {

	args := flag.Args()
	var err error = nil

	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
}
func printFile(fileName string, lineNumber bool) (err error) {

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	for i := 0; scanner.Scan(); i++ {
		if lineNumber {
			fmt.Print(i + 1)
		}
		fmt.Println(scanner.Text())
	}
	return nil

}
func main() {
	var lineNumber bool
	flag.BoolVar(&lineNumber, "n", false, "Number of line will be printed from the beginning of line")
	flag.Parse()
	filePath, err := catArgumentsValidator()
	check(err)
	err = printFile(filePath, lineNumber)
	check(err)
}
