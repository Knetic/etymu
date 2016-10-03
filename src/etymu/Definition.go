package etymu

import (
	"regexp"
)

type Definition struct {
	Patterns []*regexp.Regexp
}

func NewDefinition(patterns ...string) (*Definition, error) {

	var ret *Definition
	var regex *regexp.Regexp
	var err error

	ret = new(Definition)
	ret.Patterns = make([]*regexp.Regexp, 8)

	for _, pattern := range patterns {

		regex, err = regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}

		ret.Patterns = append(ret.Patterns, regex)
	}

	return ret, nil
}
