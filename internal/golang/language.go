package golang

import (
	"fmt"
	"unsafe"

	"grow.graphics/gd"
	classdb "grow.graphics/gd/internal/classdb"
	"grow.graphics/gd/object/Script"
	"grow.graphics/gd/object/ScriptLanguageExtension"
)

type Language struct {
	gd.Class[Language, ScriptLanguageExtension.Expert] `gd:"GoLanguage"`
}

func (lang *Language) AddGlobalConstant(name gd.StringName, value gd.Variant)      {}
func (lang *Language) AddNamedGlobalConstant(name gd.StringName, value gd.Variant) {}

func (lang *Language) AutoIndentCode(code gd.String, from, upto gd.Int) gd.String {
	return lang.Temporary.String("")
}
func (lang *Language) CanInheritFromFile() gd.Bool {
	return false
}
func (lang *Language) CompleteCode(code, path gd.String, owner gd.Object) gd.Dictionary {
	return lang.Temporary.Dictionary()
}
func (lang *Language) CreateScript() gd.Object {
	fmt.Println("CreateScript")
	return gd.Object{}
}

func (lang *Language) DebugGetCurrentStackInfo() gd.ArrayOf[gd.Dictionary] {
	return gd.NewArrayOf[gd.Dictionary](lang.Temporary)
}

func (lang *Language) DebugGetError() gd.String {
	return lang.Temporary.String("")
}

func (lang *Language) DebugGetGlobals(max_subitems, max_depth gd.Int) gd.Dictionary {
	return lang.Temporary.Dictionary()
}

func (lang *Language) DebugGetStackLevelCount() gd.Int {
	return 0
}

func (lang *Language) DebugGetStackLevelFunction(level gd.Int) gd.String {
	return lang.Temporary.String("")
}

func (lang *Language) DebugGetStackLevelInstance(level gd.Int) unsafe.Pointer {
	return nil
}

func (lang *Language) DebugGetStackLevelLine(level gd.Int) gd.Int {
	return 0
}

func (lang *Language) DebugGetStackLevelLocals(level, max_subitems, max_depth gd.Int) gd.Dictionary {
	return lang.Temporary.Dictionary()
}

func (lang *Language) DebugGetStackLevelMembers(level, max_subitems, max_depth gd.Int) gd.Dictionary {
	return lang.Temporary.Dictionary()
}

func (lang *Language) DebugParseStackLevelExpression(level gd.Int, expression gd.String, max_subitems, max_depth gd.Int) gd.String {
	return lang.Temporary.String("")
}

func (lang *Language) FindFunction(class_name, function_name gd.String) gd.Int {
	return 0
}

func (lang *Language) Finish() {}

func (lang *Language) Frame() {}

func (lang *Language) GetBuiltinTemplates(object gd.StringName) gd.ArrayOf[gd.Dictionary] {
	return gd.NewArrayOf[gd.Dictionary](lang.Temporary)
}

// GetCommentDelimeters should return the comment delimeters for the language.
func (lang *Language) GetCommentDelimeters() gd.PackedStringArray {
	return lang.Temporary.PackedStringSlice([]string{"//", "/* */"})
}

// GetDocCommentDelimiters should return the comment delimeters for the language.
func (lang *Language) GetDocCommentDelimiters() gd.PackedStringArray {
	return lang.Temporary.PackedStringSlice([]string{"//", "/* */"})
}

// GetExtension should return the file extension for source files.
func (lang *Language) GetExtension() gd.String {
	return lang.Temporary.String("go")
}

func (lang *Language) GetGlobalClassName(path gd.String) gd.Dictionary {
	return lang.Temporary.Dictionary()
}

// GetName of the languages, as it will appear in the editor.
func (lang *Language) GetName() gd.String {
	return lang.Temporary.String("Go")
}

func (lang *Language) GetPublicAnnotations() gd.ArrayOf[gd.Dictionary] {
	return gd.NewArrayOf[gd.Dictionary](lang.Temporary)
}

func (lang *Language) GetPublicConstants() gd.Dictionary {
	return lang.Temporary.Dictionary()
}

func (lang *Language) GetPublicFunctions() gd.ArrayOf[gd.Dictionary] {
	return gd.NewArrayOf[gd.Dictionary](lang.Temporary)
}

// GetRecognizedExtensions returns a list of file extensions that the language
// is aware of.
func (lang *Language) GetRecognizedExtensions() gd.PackedStringArray {
	return lang.Temporary.PackedStringSlice([]string{
		"go",
	})
}

// GetReservedWords returns a list of reserved words that cannot be used as
// identifiers.
func (lang *Language) GetReservedWords() gd.PackedStringArray {
	return lang.Temporary.PackedStringSlice([]string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	})
}

// GetStringDelimiters should return the string delimiters for the language.
func (lang *Language) GetStringDelimiters() gd.PackedStringArray {
	return lang.Temporary.PackedStringSlice([]string{"\" \"", "' '", "` `"})
}

func (lang *Language) GetType() gd.String { return lang.Temporary.String("GoScript") }

func (lang *Language) HandlesGlobalClassType(ctype gd.String) gd.Bool { return false }

func (lang *Language) HasNamedClasses() gd.Bool { return false }

func (lang *Language) Init() {}

// IsControlFlowKeyword returns true if the keyword is a control flow keyword.
func (lang *Language) IsControlFlowKeyword(keyword gd.String) gd.Bool {
	switch keyword.String() {
	case "break", "continue", "fallthrough", "return", "for", "if", "defer", "switch", "else", "goto":
		return true
	default:
		return false
	}
}

func (lang *Language) IsUsingTemplates() gd.Bool {
	return false
}

func (lang *Language) LookupCode(code, symbol, path gd.String, owner gd.Object) gd.Dictionary {
	return lang.Temporary.Dictionary()
}

func (lang *Language) MakeFunction(class_name, function_name gd.String, function_args gd.PackedStringArray) gd.String {
	return lang.Temporary.String("")
}

func (lang *Language) MakeTemplate(template, class_name, base_class_name gd.String) Script.Expert {
	fmt.Println("MakeTemplate")
	return Script.Expert{}
}

func (lang *Language) OpenInExternalEditor(script Script.Expert, line, column gd.Int) gd.Int {
	return 0
}

func (lang *Language) OverridesExternalEditor() gd.Bool {
	return false
}

func (lang *Language) ProfilingGetAccumulatedData(info_array *classdb.ScriptLanguageExtensionProfilingInfo, info_max int) int {
	return 0
}

func (lang *Language) ProfilingGetFrameData(info_array *classdb.ScriptLanguageExtensionProfilingInfo, info_max int) int {
	return 0
}

func (lang *Language) ProfilingStart()                                            {}
func (lang *Language) ProfilingStop()                                             {}
func (lang *Language) ReloadAllScripts()                                          {}
func (lang *Language) ReloadToolScript(script Script.Expert, soft_reload gd.Bool) {}
func (lang *Language) RemoveNamedGlobalConstant(name gd.StringName)               {}
func (lang *Language) SupportsBuiltinMode() gd.Bool                               { return false }
func (lang *Language) SupportsDocumentation() gd.Bool                             { return false }
func (lang *Language) ThreadEnter()                                               {}
func (lang *Language) ThreadExit()                                                {}
func (lang *Language) Validate(script, path gd.String, functions, errors, warnings, safe_lines gd.Bool) gd.Dictionary {
	return lang.Temporary.Dictionary()
}
func (lang *Language) ValidatePath(path gd.String) gd.String {
	return lang.Temporary.String("")
}
