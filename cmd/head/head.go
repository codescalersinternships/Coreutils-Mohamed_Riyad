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

func headArgumentsValidator() (string, error) {

	args := flag.Args()
	var err error = nil

	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
}

func printLinesHead(fileName string, num int) (err error) {

	f, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("failed to open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for i := 0; i < num; i++ {
		if scanner.Scan() {
			fmt.Println(scanner.Text())
		} else {
			break
		}
	}
	return nil

}

func main() {
	var num int
	flag.IntVar(&num, "n", 10, "Number of lines which will be printed from the beginning of file")
	flag.Parse()
	filePath, err := headArgumentsValidator()
	check(err)
	err = printLinesHead(filePath, num)
	check(err)
}
