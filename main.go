package main

import (
	"fmt"
	"os"
	. "etymu"
)

func main() {

	var settings RunSettings
	var lex *LexFile
	var err error

	settings, err = parseRunSettings()
	if(err != nil) {
		fatal(err, 1)
	}

	lex, err = LexFileFromPath(settings.InputPath)
	if(err != nil) {
		fatal(err, 2)
	}

	fmt.Printf("Parsed lex: %v\n", lex)
}

func fatal(fault error, code int) {

	fmt.Fprintf(os.Stderr, fault.Error())
	os.Exit(code)
}
