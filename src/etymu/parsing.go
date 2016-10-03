package etymu

import (
	"strings"
	"bufio"
	"io"
)

/*
	Parses a single "definition" line.
*/
func addDefinitionLine(lex *LexFile, line string) error {

	line = strings.TrimSpace(line)
	if(line == "" || line[0:2] == "//") {
		// comment or whitespace
		return nil
	}
	return nil
}

/*
	Parses a single "rule" line.
*/
func addRuleLine(lex *LexFile, line string) error {
	return nil
}

/*
	Reads lines from [reader], sending them one-by-one through [out], until the given [separator] is found.
	The given reader is not closed by this method, but the given channel is.
*/
func linesUntilSeparator(reader io.Reader, separator string, out chan string) error {

	var line []byte
	var bufferedReader *bufio.Reader
	var err error
	var ok bool

	defer close(out)
	
	bufferedReader = bufio.NewReader(reader)

	for {

		line, ok, err = bufferedReader.ReadLine()
		if(!ok) {
			break
		}
		if(err != nil) {
			return err
		}

		out <- string(line)
	}

	return nil
}
