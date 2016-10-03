package etymu

import (
	"regexp"
)

type Rule struct {
	Patterns 	[]*regexp.Regexp
	Action     string
}

func NewRule(action string, patterns ...*regexp.Regexp) *Rule {

	ret := new(Rule)
	ret.Patterns = patterns
	ret.Action = action
	return ret
}
