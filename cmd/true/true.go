package true

import (
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

func main() {
	var err error = nil
	args := flag.Args()

	if len(args) != 0 {
		err = fmt.Errorf("Incorrect number of arguments. Expected 0, got %d", len(args))
	}

	check(err)
	os.Exit(0)
}
