package main

import (
	"flag"
)

type RunSettings struct {

	Package string
	Language string
	OutputPath string
}

func parseRunSettings() (RunSettings, error) {

	var ret RunSettings

	flag.StringVar(&ret.Language, "l", "go", "Language to generate code for")
	flag.StringVar(&ret.Package, "p", "lexer", "Package (or module) name for generated lexer")
	flag.StringVar(&ret.OutputPath, "o", "./output.go", "Output path for generated lexer")
	flag.Parse()

	return ret, nil
}
