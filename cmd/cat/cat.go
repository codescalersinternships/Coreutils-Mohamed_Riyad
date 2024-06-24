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

func parserCat() (string, error) {

	args := flag.Args()
	var err error = nil

	if len(args) != 1 {
		err = fmt.Errorf("Incorrect number of arguments. Expected 1, got %d", len(args))
	}
	var filePath = args[0]

	return filePath, err
}
func printFile(fileName string, num int) {

	f, err := os.Open(fileName)
	check(err)

	scanner := bufio.NewScanner(f)
	check(err)
	if num == -1 {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		for i := 0; scanner.Scan() && i < num; i++ {
			fmt.Println(scanner.Text())
		}
	}
}
func main() {
	var num int
	flag.IntVar(&num, "n", -1, "Number of lines which will be printed from the beginning of file")
	flag.Parse()
	filePath, err := parserCat()
	check(err)
	printFile(filePath, num)
}
