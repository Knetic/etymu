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

	return LexFileFromStream(file)
}

func LexFileFromStream(reader io.Reader) (*LexFile, error) {

	return nil, nil
}
