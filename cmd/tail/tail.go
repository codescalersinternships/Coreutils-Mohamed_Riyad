package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parserTail() (string, error) {

	args := flag.Args()
	var err error = nil

	if len(args) != 1 {
		err = fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
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
	var end int = max(n-num, 0)
	for i := end; i < n; i++ {
		fmt.Println(lines[i])
	}

}

func main() {
	var num int
	flag.IntVar(&num, "n", 10, "Number of lines which will be printed from the beginning of file")
	flag.Parse()
	filePath, err := parserTail()
	check(err)
	printLinesTail(filePath, num)

}
