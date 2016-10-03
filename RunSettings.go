package main

import (
	"flag"
	"errors"
)

type RunSettings struct {

	Package string
	Language string
	OutputPath string
	InputPath string
}

func parseRunSettings() (RunSettings, error) {

	var ret RunSettings

	flag.StringVar(&ret.Language, "l", "go", "Language to generate code for")
	flag.StringVar(&ret.Package, "p", "lexer", "Package (or module) name for generated lexer")
	flag.StringVar(&ret.OutputPath, "o", "./output.go", "Output path for generated lexer")
	flag.Parse()

	ret.InputPath = flag.Arg(0)
	if(ret.InputPath == "") {
		return ret, errors.New("First positional parameter must be an input file")
	}

	return ret, nil
}
