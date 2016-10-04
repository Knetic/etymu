package etymu

func GenerateGo(file *LexFile, module string, out chan []byte) {

	buffer := NewBufferedFormatString("\t")

	buffer.Printfln("package %s\n", module)
	buffer.Printfln("import \"regexp\"\n")

	generateGoTypes(file, buffer)
	generateGoRules(file, buffer)
	generateGoLexer(file, buffer)

	out <- []byte(buffer.String())
}

func generateGoTypes(file *LexFile, buffer *BufferedFormatString) {

	// generate token constants
	buffer.Printfln("type TokenKind uint\n")
	buffer.Printfln("const UNKNOWN TokenKind = 0")

	for index, action := range file.GetAllActionNames() {
		buffer.Printfln("const %s TokenKind = %d", action, index+1)
	}
	buffer.Printfln("")

	writeStruct(buffer, "Token", "value string", "kind TokenKind")
	writeStruct(buffer, "lexerRule", "pattern *regexp.Regexp", "kind TokenKind")
}

func generateGoRules(file *LexFile, buffer *BufferedFormatString) {

	buffer.Printfln("var lexerRules []lexerRule = []lexerRule {")
	buffer.AddIndentation(1)

	for _, rule := range file.Rules {
		for _, pattern := range rule.Patterns {
			buffer.Printfln("{")
			buffer.AddIndentation(1)
			buffer.Printfln("pattern: regexp.MustCompile(\"%s\"),", escapeGoPattern(pattern))

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
