package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	var withoutNewLine bool
	flag.BoolVar(&withoutNewLine, "n", false, "print without newline in the end of string")
	flag.Parse()
	args := flag.Args()
	text := strings.Join(args, " ")
	if !withoutNewLine {
		fmt.Println(text)
	} else {
		fmt.Print(text)
	}
}
