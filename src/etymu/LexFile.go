package etymu

import (
	"path/filepath"
	"os"
	"io"
)

type LexFile struct {

	Definitions []Definition
	Rules []Rule
}

func LexFileFromPath(path string) (*LexFile, error) {

	var file *os.File
	var err error

	path, err = filepath.Abs(path)
	if(err != nil) {
		return nil, err
	}

	file, err = os.Open(path)
	if(err != nil) {
		return nil, err
	}
	defer file.Close()

	return LexFileFromStream(file)
}

func LexFileFromStream(reader io.Reader) (*LexFile, error) {

	var lines chan string
	var readErr, err error

	ret := new(LexFile)

	// see parsing.go for the bulk of the logic for this.
	// read definitions
	lines = make(chan string)
	go func(){readErr = linesUntilSeparator(reader, "%%", lines)}()
	for line := range lines {

		err = addDefinitionLine(ret, line)
		if(err != nil) {
			return nil, err
		}
	}

	if(readErr != nil) {
		return nil, readErr
	}

	// read rules
	lines = make(chan string)
	go func(){readErr = linesUntilSeparator(reader, "%%", lines)}()
	for line := range lines {

		err = addRuleLine(ret, line)
		if(err != nil) {
			return nil, err
		}
	}

	return ret, nil
}
