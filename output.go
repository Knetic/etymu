package main

import (
	"errors"
	. "etymu"
	"fmt"
	"os"
	"path/filepath"
)

/*
	Generates code for the given [file],
	outputting the unicode representation of it to [out].
*/
type Codegen func(file *LexFile, out chan []byte)

func Generate(language string, path string, lex *LexFile) error {

	var generator Codegen
	var writeErr error

	switch language {
	case "go":
		generator = GenerateGo
	default:
		errorMsg := fmt.Sprintf("Unable to generate code for unknown language '%s'\n", language)
		return errors.New(errorMsg)
	}

	generated := make(chan []byte, 8)

	go func() { writeErr = writeOutputPath(path, generated) }()
	generator(lex, generated)

	if writeErr != nil {
		return writeErr
	}
	return nil
}

/*
	Writes all the output from the given [in] channel to a file at the given [path].
	Returns any errors it found while doing so.
	The channel is not closed by this method.
*/
func writeOutputPath(path string, in <-chan []byte) error {

	var file *os.File
	var err error

	path, err = filepath.Abs(path)
	if err != nil {
		return err
	}

	file, err = os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	for buffered := range in {
		file.Write(buffered)
	}

	return nil
}
