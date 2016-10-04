package main

import (
	. "etymu"
	"fmt"
	"os"
)

func main() {

	var settings RunSettings
	var lex *LexFile
	var err error

	settings, err = parseRunSettings()
	if err != nil {
		fatal(err, 1)
	}

	lex, err = LexFileFromPath(settings.InputPath)
	if err != nil {
		fatal(err, 2)
	}

	err = Generate(settings.Language, settings.Package, settings.OutputPath, lex)
	if err != nil {
		fatal(err, 3)
	}
}

func fatal(fault error, code int) {

	fmt.Fprintf(os.Stderr, "%v\n", fault.Error())
	os.Exit(code)
}
