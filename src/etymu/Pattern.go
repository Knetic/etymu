package etymu

import (
	"regexp"
	"strings"
)

type Pattern struct {
	value   string
	isRegex bool
}

func parseRulePattern(pattern string) (Pattern, error) {

	var err error

	if strings.HasPrefix(pattern, "\"") && strings.HasSuffix(pattern, "\"") {
		return Pattern{
			value:   pattern[1 : len(pattern)-1],
			isRegex: false,
		}, nil
	}

	_, err = regexp.Compile(pattern)
	if err != nil {
		return Pattern{}, err
	}

	return Pattern{
		value:   pattern,
		isRegex: true,
	}, nil
}
