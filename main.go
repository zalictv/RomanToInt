package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		fmt.Println(RomanToInt(args[1]))

	} else {
		fmt.Println("Usage: ", args[0], " ROMAN-NUMBER")
	}
}
