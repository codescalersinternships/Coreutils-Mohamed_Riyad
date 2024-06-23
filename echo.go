package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func parserAndPrint(input string) {
	words := strings.Fields(input)
	var err error
	if len(words) == 0 || words[0] != "echo" {
		err = fmt.Errorf("error: command must start with 'echo'")
	}
	if len(words) == 3 {
		fmt.Print(words[2])
	} else if len(words) == 2 {
		fmt.Println(words[1])
	} else {
		err = fmt.Errorf("error: incorrect number of arguments")
	}
	check(err)

}
func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	input, err := reader.ReadString('\n')
	check(err)
	parserAndPrint(input)

}
