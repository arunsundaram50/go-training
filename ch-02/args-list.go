package main

import (
	"fmt"
	"os"
)

// program to print input arguments
func main() {
	// get input arguments
	args := os.Args[1:]
	// range over each input argument and print it
	for _, arg := range args {
		fmt.Println(arg)
	}
}
