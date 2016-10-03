package etymu

import (
	"bufio"
	"io"
	"strings"
	"errors"
	"unicode"
)

/*
	Parses a single "definition" line.
*/
func addDefinitionLine(lex *LexFile, line string) error {

	var name, pattern string
	var err error

	name, pattern, err = getWhitespaceDelimitedString(line)
	if(err != nil) {
		return err
	}
	if(name == "" && pattern == "") {
		return nil
	}

	lex.AddDefinition(name, pattern)
	return nil
}

/*
	Parses a single "rule" line.
*/
func addRuleLine(lex *LexFile, line string) error {

	var left, right string
	var err error

	left, right, err = getWhitespaceDelimitedString(line)
	if(err != nil) {
		return err
	}
	if(left == "" && right == "") {
		return nil
	}

	// get a set of patterns from the left side

	// get the action (in braces) from the right side

	return nil
}

func getWhitespaceDelimitedString(line string) (string, string, error) {

	var leftIdx, rightIdx int
	var lineLen int

	line = strings.TrimSpace(line)
	if line == "" || line[0:2] == "//" {
		// comment or whitespace
		return "", "", nil
	}

	lineLen = len(line)

	for leftIdx = 0; leftIdx < lineLen; leftIdx++ {
		if(unicode.IsSpace(rune(line[leftIdx]))) {
			break
		}
	}

	for rightIdx = leftIdx+1; rightIdx < lineLen; rightIdx++ {
		if(!unicode.IsSpace(rune(line[rightIdx]))) {
			break
		}
	}

	if(leftIdx >= lineLen || rightIdx >= lineLen) {
		return "", "", errors.New("Unable to parse definition, did not contain a whitespace-separated name and pattern.")
	}

	return line[0:leftIdx], line[rightIdx:], nil
}

/*
	Reads lines from [reader], sending them one-by-one through [out], until the given [separator] is found.
	The given reader is not closed by this method, but the given channel is.
*/
func linesUntilSeparator(reader io.Reader, separator string, out chan string) error {

	var line string
	var bufferedReader *bufio.Reader
	var err error

	defer close(out)

	bufferedReader = bufio.NewReader(reader)

	for {

		line, err = bufferedReader.ReadString('\n')
		if err != nil {
			if(err == io.EOF) {
				break
			}
			return err
		}

		if(strings.HasPrefix(line, separator)) {
			break
		}

		out <- line
	}

	return nil
}
