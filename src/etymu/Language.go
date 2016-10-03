package etymu

type Language uint16

const (
	LANG_UNKNOWN Language = iota
	LANG_GO
)

// Map of all valid input language names and their ordinal.
var LanguageNameMap map[string]Language = map[string]Language{
	"go": LANG_GO,
}

// Map of all known languages and their sourcecode file extension.
var LanguageExtensionMap map[Language]string = map[Language]string{
	LANG_GO: "go",
}
