package golang

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Script"
	"graphics.gd/classdb/ScriptLanguage"
	"graphics.gd/classdb/ScriptLanguageExtension"
	"graphics.gd/internal/gdextension"
	"graphics.gd/variant/Object"
)

type Language struct {
	ScriptLanguageExtension.Extension[Language] `gd:"GoLanguage"`
}

func (lang *Language) GetName() string { return "Go" }
func (lang *Language) Init()           {}
func (lang *Language) GetType() string { return classdb.NameFor[GoScript]() }

// GetExtension should return the file extension for source files.
func (lang *Language) GetExtension() string { return "go" }

func (lang *Language) Finish() {}

// GetReservedWords returns a list of reserved words that cannot be used as
// identifiers.
func (lang *Language) GetReservedWords() []string {
	return []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}
}

// IsControlFlowKeyword returns true if the keyword is a control flow keyword.
func (lang *Language) IsControlFlowKeyword(keyword string) bool {
	switch keyword {
	case "break", "continue", "fallthrough", "return", "for", "if", "defer", "switch", "else", "goto":
		return true
	default:
		return false
	}
}

func (lang *Language) GetCommentDelimiters() []string {
	return []string{"//", "/* "}
}

// GetDocCommentDelimiters should return the comment delimeters for the language.
func (lang *Language) GetDocCommentDelimiters() []string {
	return []string{"//", "/* "}
}

// GetStringDelimiters should return the string delimiters for the language.
func (lang *Language) GetStringDelimiters() []string {
	return []string{"\" \"", "' '", "` `"}
}

func (lang *Language) MakeTemplate(template string, class_name string, base_class_name string) Script.Instance {
	return Script.Nil
}

func (lang *Language) GetBuiltinTemplates(object string) [][]ScriptLanguageExtension.Template {
	return nil
}

func (lang *Language) IsUsingTemplates() bool {
	return false
}

func (lang *Language) Validate(script, path string, functions, errors, warnings, safe_lines bool) ScriptLanguageExtension.Validation {
	return ScriptLanguageExtension.Validation{}
}
func (lang *Language) ValidatePath(path string) string {
	return ""
}
func (lang *Language) CreateScript() Object.Instance {
	return Object.Nil
}

func (lang *Language) HasNamedClasses() bool       { return false }
func (lang *Language) SupportsBuiltinMode() bool   { return false }
func (lang *Language) SupportsDocumentation() bool { return false }
func (lang *Language) CanInheritFromFile() bool    { return false }

func (lang *Language) FindFunction(class_name, function_name string) int {
	return 0
}
func (lang *Language) MakeFunction(class_name, function_name string, function_args []string) string {
	return ""
}
func (lang *Language) CanMakeFunction() bool {
	return false
}
func (lang *Language) OpenInExternalEditor(script Script.Instance, line int, column int) error {
	return nil
}
func (lang *Language) OverridesExternalEditor() bool {
	return false
}
func (lang *Language) PreferredFileNameCasing() ScriptLanguage.ScriptNameCasing {
	return ScriptLanguage.ScriptNameCasingSnakeCase
}
func (lang *Language) CompleteCode(code string, path string, owner Object.Instance) ScriptLanguageExtension.Completion {
	return ScriptLanguageExtension.Completion{}
}
func (lang *Language) LookupCode(code, symbol, path string, owner Object.Instance) ScriptLanguageExtension.Code {
	return ScriptLanguageExtension.Code{}
}
func (lang *Language) AutoIndentCode(code string, from, upto int) string {
	return ""
}
func (lang *Language) AddGlobalConstant(name string, value any)      {}
func (lang *Language) AddNamedGlobalConstant(name string, value any) {}
func (lang *Language) RemoveNamedGlobalConstant(name string)         {}

func (lang *Language) ThreadEnter() {}
func (lang *Language) ThreadExit()  {}

func (lang *Language) DebugGetError() string {
	return ""
}
func (lang *Language) DebugGetStackLevelCount() int {
	return 0
}
func (lang *Language) DebugGetStackLevelLine(level int) int {
	return 0
}
func (lang *Language) DebugGetStackLevelFunction(level int) string {
	return ""
}

func (lang *Language) DebugGetStackLevelSource(level int) string {
	return ""
}

func (lang *Language) DebugGetStackLevelLocals(level, max_subitems, max_depth int) ScriptLanguageExtension.StackLevelLocals {
	return ScriptLanguageExtension.StackLevelLocals{}
}

func (lang *Language) DebugGetStackLevelMembers(level, max_subitems, max_depth int) ScriptLanguageExtension.StackLevelMembers {
	return ScriptLanguageExtension.StackLevelMembers{}
}

func (lang *Language) DebugGetStackLevelInstance(level int) gdextension.Pointer {
	return 0
}

func (lang *Language) DebugGetGlobals(max_subitems, max_depth int) ScriptLanguageExtension.Globals {
	return ScriptLanguageExtension.Globals{}
}

func (lang *Language) DebugParseStackLevelExpression(level int, expression string, max_subitems, max_depth int) string {
	return ""
}

func (lang *Language) DebugGetCurrentStackInfo() []ScriptLanguageExtension.StackInfo {
	return nil
}

func (lang *Language) ReloadAllScripts()
func (lang *Language) ReloadScripts(scripts []Script.Instance, soft_reload bool)  {}
func (lang *Language) ReloadToolScript(scripts Script.Instance, soft_reload bool) {}

// GetRecognizedExtensions returns a list of file extensions that the language
// is aware of.
func (lang *Language) GetRecognizedExtensions() []string {
	return []string{"go"}
}

func (lang *Language) GetPublicFunctions() [][]struct{} {
	return nil
}

func (lang *Language) GetPublicConstants() []ScriptLanguageExtension.Constant {
	return nil
}

func (lang *Language) GetPublicAnnotations() [][]struct{} {
	return nil
}

func (lang *Language) ProfilingStart() {}
func (lang *Language) ProfilingStop()  {}

func (lang *Language) ProfilingSetSaveNativeCalls(enable bool) {}

func (lang *Language) ProfilingGetAccumulatedData(info_array *ScriptLanguageExtension.ProfilingInfo, info_max int) int {
	return 0
}

func (lang *Language) ProfilingGetFrameData(info_array *ScriptLanguageExtension.ProfilingInfo, info_max int) int {
	return 0
}

func (lang *Language) Frame() {}

func (lang *Language) HandlesGlobalClassType(ctype string) bool { return false }

func (lang *Language) GetGlobalClassName(path string) ScriptLanguageExtension.ClassName {
	return ScriptLanguageExtension.ClassName{}
}
