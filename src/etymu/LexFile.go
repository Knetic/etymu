package etymu

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type LexFile struct {
	Definitions map[string]*regexp.Regexp
	Rules       []Rule
}

func LexFileFromPath(path string) (*LexFile, error) {

	var file *os.File
	var err error

	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	file, err = os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return LexFileFromStream(file)
}

func LexFileFromStream(reader io.Reader) (*LexFile, error) {

	var lines chan string
	var readErr, err error

	ret := new(LexFile)
	ret.Definitions = make(map[string]*regexp.Regexp)

	// see parsing.go for the bulk of the logic for this.
	// read definitions
	lines = make(chan string)
	go func() { readErr = linesUntilSeparator(reader, "%%", lines) }()
	for line := range lines {

		err = addDefinitionLine(ret, line)
		if err != nil {
			return nil, err
		}
	}

	if readErr != nil {
		return nil, readErr
	}

	// read rules
	lines = make(chan string)
	go func() { readErr = linesUntilSeparator(reader, "%%", lines) }()
	for line := range lines {

		err = addRuleLine(ret, line)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (this *LexFile) AddDefinition(name string, pattern string) error {

	regex, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	this.Definitions[name] = regex
	return err
}
