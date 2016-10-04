package main

import (
	"errors"
	. "etymu"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

/*
	Generates code for the given [file],
	outputting the unicode representation of it to [out].
*/
type Codegen func(file *LexFile, module string, out chan []byte)

func Generate(language Language, module string, path string, lex *LexFile) error {

	var generator Codegen
	var file *os.File
	var err error

	switch language {
	case LANG_GO:
		generator = GenerateGo
	default:
		errorMsg := fmt.Sprintf("Program incomplete, implementation not found for language '%v' ('%s'). This is a problem for the developer.\n", language, LanguageExtensionMap[language])
		return errors.New(errorMsg)
	}

	file, err = createOutputPath(path)
	if err != nil {
		return err
	}

	generated := make(chan []byte)

	go func() {
		defer close(generated)
		generator(lex, module, generated)
	}()

	writeOutputPath(file, generated)
	if err != nil {
		return err
	}
	return nil
}

func findOutputName(inputPath string) string {
	return filepath.Base(inputPath)
}

func createOutputPath(path string) (*os.File, error) {

	var err error

	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	return os.Create(path)
}

/*
	Writes all the output from the given [in] channel to a file at the given [path].
	Returns any errors it found while doing so.
	The channel is not closed by this method.
*/
func writeOutputPath(writer io.Writer, in <-chan []byte) {
	for buffered := range in {
		writer.Write(buffered)
	}
}
