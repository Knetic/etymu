package etymu

func GenerateGo(file *LexFile, module string, out chan []byte) {

	buffer := NewBufferedFormatString("\t")

	buffer.Printfln("package %s\n", module)
	generateGoActions(file, buffer)

	out <- []byte(buffer.String())
}

func generateGoActions(file *LexFile, buffer *BufferedFormatString) {

	for index, action := range file.GetAllActionNames() {
		buffer.Printfln("const %s uint = %d", action, index)
	}
}
