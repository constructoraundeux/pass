package main

import (
	"fmt"
	"os"
)

// safe wraps any function that returns an error,
// if an error is returned, it prints it and exits
// with a non-zero exit code.
func safe(err error) {
	if err != nil {
		out(fmt.Sprintf("Tenemos problemas: %s", err.Error()))
	}
}

// info prints a message to os.Stdout
func info(msg string) {
	fmt.Fprintf(os.Stdout, msg+"\n")
}

// out prints a message to os.Stderr and exits
// with a non-zero exit code.
func out(msg string) {
	fmt.Fprintf(os.Stderr, msg+"\n")
	os.Exit(1)
}
