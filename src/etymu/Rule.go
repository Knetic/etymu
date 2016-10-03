package etymu

type Rule struct {

	Definition Definition
	Action string
}

func NewRule(definition Definition, action string) *Rule {

	ret := new(Rule)
	ret.Definition = definition
	ret.Action = action

	return ret
}
