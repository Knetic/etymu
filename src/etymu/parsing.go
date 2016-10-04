package etymu

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"unicode"
	"fmt"
)

/*
	Parses a single "definition" line.
*/
func addDefinitionLine(lex *LexFile, line string) error {

	var name, pattern string
	var err error

	name, pattern, err = getWhitespaceDelimitedString(line)
	if err != nil {
		return err
	}
	if name != "" && pattern == "" {
		return errors.New("Unable to parse definition, did not contain a whitespace-separated name and pattern.")
	}
	if name == "" && pattern == "" {
		return nil
	}

	lex.AddDefinition(name, pattern)
	return nil
}

/*
	Parses a single "rule" line.
*/
func addRuleLine(lex *LexFile, line string) error {

	var patterns []string
	var left, right string
	var err error

	left, right, err = getWhitespaceDelimitedString(line)
	if err != nil {
		return err
	}
	if left == "" && right == "" {
		return nil
	}

	// get a set of patterns from the left side
	patterns, err = lex.resolvePatterns(strings.Split(left, "|")...)
	if err != nil {
		return err
	}

	if(len(right) > 0) {

		if(right[0:1] != "{") {
			errorMsg := fmt.Sprintf("No opening brace to action ('%s')", right)
			return errors.New(errorMsg)
		}

		closeIdx := strings.Index(right, "}")
		if(closeIdx <= -1) {
			errorMsg := fmt.Sprintf("No closing brace to action ('%s')", right)
			return errors.New(errorMsg)
		}

		right = right[1:closeIdx]
	}

	lex.AddRule(right, patterns...)
	return nil
}

func getWhitespaceDelimitedString(line string) (string, string, error) {

	var leftIdx, rightIdx int
	var lineLen int

	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "//") {
		// comment or whitespace
		return "", "", nil
	}

	lineLen = len(line)

	for leftIdx = 0; leftIdx < lineLen; leftIdx++ {
		if unicode.IsSpace(rune(line[leftIdx])) {
			break
		}
	}

	if leftIdx < lineLen-1 {
		rightIdx = leftIdx + 1
	} else {
		rightIdx = leftIdx
	}

	for ; rightIdx < lineLen; rightIdx++ {
		if !unicode.IsSpace(rune(line[rightIdx])) {
			break
		}
	}

	return line[0:leftIdx], line[rightIdx:], nil
}

/*
	Reads lines from [reader], sending them one-by-one through [out], until the given [separator] is found.
	The given reader is not closed by this method, but the given channel is.
*/
func linesUntilSeparator(reader *bufio.Reader, separator string, out chan string) error {

	var line string
	var err error

	defer close(out)

	for {

		line, err = reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if strings.HasPrefix(line, separator) {
			break
		}

		out <- line
	}

	return nil
}
