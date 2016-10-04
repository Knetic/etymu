package etymu

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type LexFile struct {
	Definitions map[string]string
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
	var bufferedReader *bufio.Reader
	var readErr, err error

	ret := new(LexFile)
	ret.Definitions = make(map[string]string)

	bufferedReader = bufio.NewReader(reader)

	// see parsing.go for the bulk of the logic for this.
	// read definitions
	lines = make(chan string)
	go func() { readErr = linesUntilSeparator(bufferedReader, "%%", lines) }()
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
	go func() { readErr = linesUntilSeparator(bufferedReader, "%%", lines) }()
	for line := range lines {

		err = addRuleLine(ret, line)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
}

func (this *LexFile) AddDefinition(name string, pattern string) error {
	this.Definitions[name] = pattern
	return nil
}

func (this *LexFile) AddRule(action string, patterns ...string) error {

	var rulePatterns []Pattern
	var rulePattern Pattern
	var err error

	for _, pattern := range patterns {

		rulePattern, err = parseRulePattern(pattern)
		if err != nil {
			return err
		}

		rulePatterns = append(rulePatterns, rulePattern)
	}

	rule := Rule{
		Action:   strings.TrimSpace(action),
		Patterns: rulePatterns,
	}

	this.Rules = append(this.Rules, rule)
	return nil
}

func (this *LexFile) GetAllActionNames() []string {

	var ret []string
	var found bool

	for _, rule := range this.Rules {

		// skip empties
		if rule.Action == "" {
			continue
		}

		// see if it's already listed
		found = false
		for _, extant := range ret {
			if rule.Action == extant {
				found = true
				break
			}
		}

		if found {
			continue
		}
		ret = append(ret, rule.Action)
	}
	return ret
}

/*
	Takes all the given patterns and replaces any definitions with the actual pattern.
*/
func (this *LexFile) resolvePatterns(patterns ...string) ([]string, error) {

	var ret []string
	var strippedPattern string

	for _, pattern := range patterns {

		strippedPattern = pattern[1 : len(pattern)-1]

		if pattern[0] != '{' {
			ret = append(ret, strippedPattern)
			continue
		}

		resolvedPattern, found := this.Definitions[strippedPattern]
		if !found {
			errorMsg := fmt.Sprintf("Unable to find a definition for '%s'", strippedPattern)
			return ret, errors.New(errorMsg)
		}

		ret = append(ret, resolvedPattern)
	}

	return ret, nil
}
