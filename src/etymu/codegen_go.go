package etymu

func GenerateGo(file *LexFile, module string, out chan []byte) {

	buffer := NewBufferedFormatString("\t")

	buffer.Printfln("package %s\n", module)
	generateGoTypes(file, buffer)

	out <- []byte(buffer.String())
}

func generateGoTypes(file *LexFile, buffer *BufferedFormatString) {

	// generate token constants
	buffer.Printfln("type TokenKind uint")
	buffer.Printfln("const UNKNOWN Token = 0")

	for index, action := range file.GetAllActionNames() {
		buffer.Printfln("const %s Token = %d", action, index+1)
	}
	buffer.Printfln("")

	writeStruct(buffer, "Token", "value string", "kind TokenKind")
	writeStruct(buffer, "lexerRule", "pattern string", "kind TokenKind")
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
