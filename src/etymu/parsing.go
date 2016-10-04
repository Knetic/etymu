package etymu

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode"
)

/*
	Parses a single "definition" line.
*/
func addDefinitionLine(lex *LexFile, line string) error {

	var name, right string
	var pattern Pattern
	var err error

	name, right, err = getWhitespaceDelimitedString(line)
	if err != nil {
		return err
	}
	if name != "" && right == "" {
		return errors.New("Unable to parse definition, did not contain a whitespace-separated name and pattern.")
	}
	if name == "" && right == "" {
		return nil
	}

	pattern, err = parseRulePattern(right)
	if err != nil {
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
	var resolvedPatterns []Pattern
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
	patterns, err = parseRulePatterns(left)
	if err != nil {
		return err
	}

	resolvedPatterns, err = lex.resolvePatterns(patterns...)
	if err != nil {
		return err
	}

	if len(right) > 0 {

		if right[0:1] != "{" {
			errorMsg := fmt.Sprintf("No opening brace to action ('%s')", right)
			return errors.New(errorMsg)
		}

		closeIdx := strings.Index(right, "}")
		if closeIdx <= -1 {
			errorMsg := fmt.Sprintf("No closing brace to action ('%s')", right)
			return errors.New(errorMsg)
		}

		right = right[1:closeIdx]
	}

	lex.AddRule(right, resolvedPatterns...)
	return nil
}

func parseRulePatterns(rule string) ([]string, error) {

	var ret []string
	var pattern string
	var char rune
	var ok bool

	runeChan := make(chan rune)
	go readRunes(rule, runeChan)

	for char = range runeChan {

		if char == '{' {
			pattern, ok = readChanUntil(runeChan, '}')
		}
		if char == '\'' {
			pattern, ok = readChanUntil(runeChan, '\'')
		}
		if char == '"' {
			pattern, ok = readChanUntil(runeChan, '"')
		}

		if !ok {
			errorMsg := fmt.Sprintf("No corrosponding close character found for '%s'", string(char))
			return ret, errors.New(errorMsg)
		}

		pattern = string(char) + pattern

		ret = append(ret, pattern)

		char, ok = <-runeChan
		if !ok {
			// end of input
			break
		}

		if char != '|' {
			errorMsg := fmt.Sprintf("No pipe separator between rule patterns (found '%s')", string(char))
			return ret, errors.New(errorMsg)
		}
	}

	return ret, nil
}

func readRunes(input string, out chan rune) {
	for _, char := range input {
		out <- char
	}
	close(out)
}

func readChanUntil(in <-chan rune, delim rune) (string, bool) {

	var ret string
	var escaped bool

	escaped = false
	for char := range in {

		ret = ret + string(char)

		if escaped {
			escaped = false
			continue
		}

		if char == delim {
			return ret, true
		}
		escaped = (char == '\\')
	}
	return ret, false
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
