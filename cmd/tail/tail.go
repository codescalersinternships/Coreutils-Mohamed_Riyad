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

func tailArgumentsValidator() (string, error) {

	args := flag.Args()
	var err error = nil

	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
}

func printLinesTail(fileName string, num int) (err error) {

	f, err := os.Open(fileName)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = lines[0 : len(lines)-1]
	n := len(lines)
	end := max(n-num, 0)
	for i := end; i < n; i++ {
		fmt.Println(lines[i])
	}
	return nil
}

func main() {
	var num int
	flag.IntVar(&num, "n", 10, "Number of lines which will be printed from the beginning of file")
	flag.Parse()
	filePath, err := tailArgumentsValidator()
	check(err)
	err = printLinesTail(filePath, num)
	check(err)

}
