package main

import (
	"errors"
	. "etymu"
	"flag"
	"fmt"
	"strings"
)

type RunSettings struct {
	Package    string
	Language   Language
	OutputPath string
	InputPath  string
}

func parseRunSettings() (RunSettings, error) {

	var ret RunSettings
	var language string
	var ok bool

	flag.StringVar(&language, "l", "go", "Language to generate code for")
	flag.StringVar(&ret.Package, "p", "lexer", "Package (or module) name for generated lexer")
	flag.StringVar(&ret.OutputPath, "o", "", "Output path for generated lexer")
	flag.Parse()

	ret.Language, ok = LanguageNameMap[language]
	if !ok {
		errorMsg := fmt.Sprintf("Language '%s' not recognized", language)
		return ret, errors.New(errorMsg)
	}

	ret.InputPath = flag.Arg(0)
	if ret.InputPath == "" {
		return ret, errors.New("First positional parameter must be an input file")
	}

	// if no output is specified, use the input filename as the basename for the generated file.
	if ret.OutputPath == "" {
		ret.OutputPath = fmt.Sprintf("./%s", findOutputName(ret.InputPath))
	}

	ret.OutputPath = ensureExtension(ret.OutputPath, ret.Language)
	return ret, nil
}

func ensureExtension(path string, language Language) string {

	extension := LanguageExtensionMap[language]
	if strings.HasSuffix(path, extension) {
		return path
	}
	return path + "." + extension
}
