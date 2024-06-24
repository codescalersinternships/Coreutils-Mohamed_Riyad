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

func parserHead() (string, error) {

	args := flag.Args()
	var err error = nil

	if len(args) != 1 {
		err = fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
}

func printLinesHead(fileName string, num int) {

	f, err := os.Open(fileName)
	if err != nil {
		check(fmt.Errorf("failed to open file"))
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

	if err := scanner.Err(); err != nil {
		check(fmt.Errorf("error while reading file"))
	}
}

func main() {
	var num int
	flag.IntVar(&num, "n", 10, "Number of lines which will be printed from the beginning of file")
	flag.Parse()
	filePath, err := parserHead()
	check(err)
	printLinesHead(filePath, num)

}