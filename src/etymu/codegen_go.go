package etymu

func GenerateGo(file *LexFile, module string, out chan []byte) {

	buffer := NewBufferedFormatString("\t")

	buffer.Printfln("package %s\n", module)

	generateGoImports(file, buffer)
	generateGoTypes(file, buffer)
	generateLexerFunctions(file, buffer)
	generateGoRules(file, buffer)
	generateGoLexer(file, buffer)

	out <- []byte(buffer.String())
}

func generateGoImports(file *LexFile, buffer *BufferedFormatString) {
	buffer.Printfln("import (")
	buffer.AddIndentation(1)
	buffer.Printfln("\"fmt\"")
	buffer.Printfln("\"regexp\"")
	buffer.Printfln("\"strings\"")
	buffer.Printfln("\"errors\"")
	buffer.AddIndentation(-1)
	buffer.Printfln(")\n")
}

func generateGoTypes(file *LexFile, buffer *BufferedFormatString) {

	// generate token constants
	buffer.Printfln("type TokenKind uint\n")
	buffer.Printfln("const UNKNOWN TokenKind = 0")

	for index, action := range file.GetAllActionNames() {
		buffer.Printfln("const %s TokenKind = %d", action, index+1)
	}
	buffer.Printfln("")

	writeStruct(buffer, "Token", "Value string", "Kind TokenKind")
	writeStruct(buffer, "lexerRule", "pattern *regexp.Regexp", "match string", "kind TokenKind")
}

func generateLexerFunctions(file *LexFile, buffer *BufferedFormatString) {

	// TODO: maybe make two classes implementing a common interface, to save runtime jumps?
	buffer.Printfln("func (this lexerRule) applies(input string) bool {")
	buffer.AddIndentation(1)

	buffer.Printfln("if(this.pattern == nil) {")
	buffer.AddIndentation(1)
	buffer.Printfln("return strings.HasPrefix(input, this.match)")
	buffer.AddIndentation(-1)
	buffer.Printfln("} else {")
	buffer.AddIndentation(1)
	buffer.Printfln("return this.pattern.Match([]byte(input))")
	buffer.AddIndentation(-1)
	buffer.Printfln("}")

	buffer.AddIndentation(-1)
	buffer.Printfln("}\n")
}

func generateGoRules(file *LexFile, buffer *BufferedFormatString) {

	buffer.Printfln("var lexerRules []lexerRule = []lexerRule {")
	buffer.AddIndentation(1)

	for _, rule := range file.Rules {
		for _, pattern := range rule.Patterns {
			buffer.Printfln("{")
			buffer.AddIndentation(1)

			if pattern.isRegex {
				buffer.Printfln("pattern: regexp.MustCompile(\"%s\"),", escapeGoPattern(pattern))
			} else {
				buffer.Printfln("match: \"%s\",", escapeGoPattern(pattern))
			}

			if rule.Action == "" {
				buffer.Printfln("kind: UNKNOWN,")
			} else {
				buffer.Printfln("kind: %s,", rule.Action)
			}
			buffer.AddIndentation(-1)
			buffer.Printfln("},")
		}
	}

	buffer.AddIndentation(-1)
	buffer.Printfln("}\n")
}

func generateGoLexer(file *LexFile, buffer *BufferedFormatString) {

	buffer.Printfln("func Lex(input string) ([]Token, error) {")
	buffer.AddIndentation(1)
	buffer.Printfln("var ret []Token")
	buffer.Printfln("var matchedRule lexerRule")
	buffer.Printfln("var t string")
	buffer.Printfln("var ruleMatches uint16")

	// TODO: iterate character by character, adding to a buffer, until we detect that only one rule matches the given buffer
	// if more than one rule always applies, return an ambiguity error
	// double TODO: ambiguity errors suck, those ideally should be detectable during codegen-time

	buffer.Printfln("for _, char := range input {")
	buffer.AddIndentation(1)
	buffer.Printfln("t += string(char)")
	buffer.Printfln("ruleMatches = 0")

	// check every rule to see if this matches exactly one
	buffer.Printfln("for _, rule := range lexerRules {")
	buffer.AddIndentation(1)
	buffer.Printfln("if rule.applies(t) {")
	buffer.AddIndentation(1)
	buffer.Printfln("matchedRule = rule")
	buffer.Printfln("ruleMatches++")
	buffer.Printfln("if ruleMatches >= 2 {break}")
	buffer.AddIndentation(-1)
	buffer.Printfln("}\n")

	buffer.AddIndentation(-1) // for
	buffer.Printfln("}\n")

	// if more than one rule matches, keep adding to the string until we have only one.
	buffer.Printfln("if ruleMatches > 1 {continue}")
	buffer.Printfln("token := Token{Kind: matchedRule.kind, Value: t}")
	buffer.Printfln("ret = append(ret, token)")
	buffer.Printfln("t = \"\"")

	buffer.AddIndentation(-1)
	buffer.Printfln("}\n")

	buffer.Printfln("if len(t) > 0 {")
	buffer.AddIndentation(1)
	buffer.Printfln("errorMsg := fmt.Sprintf(\"Hanging token which matched more than one lexer rule: '%%s'\", t)")
	buffer.Printfln("return ret, errors.New(errorMsg)")
	buffer.AddIndentation(-1)
	buffer.Printfln("}\n")

	buffer.Printfln("return ret, nil")
	buffer.AddIndentation(-1)
	buffer.Printfln("}")
}

func escapeGoPattern(input Pattern) string {

	var val string

	val = input.value
	//val = strings.Replace(val, "\"", "\\\"", -1)
	//val = strings.Replace(val, "\\", "\\\\", -1)
	return val
}

func writeStruct(buffer *BufferedFormatString, name string, fields ...string) {

	buffer.Printfln("type %s struct {", name)
	buffer.AddIndentation(1)

	for _, field := range fields {
		buffer.Printfln("%s", field)
	}

	buffer.AddIndentation(-1)
	buffer.Printfln("}\n")
}
