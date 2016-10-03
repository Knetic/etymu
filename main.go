package main

import (
	"fmt"
	"os"
)

func main() {

	var settings RunSettings
	var err error

	settings, err = parseRunSettings()
	if(err != nil) {
		fatal(err, 1)
	}

	fmt.Printf("Hello, etymu\n")
	fmt.Printf("Settings: %v\n", settings)
}

func fatal(fault error, code int) {

	fmt.Fprintf(os.Stderr, fault.Error())
	os.Exit(code)
}
