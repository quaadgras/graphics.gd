package golang

import (
	"graphics.gd/classdb/Script"
	"graphics.gd/classdb/ScriptExtension"
	"graphics.gd/classdb/ScriptLanguage"
	"graphics.gd/internal/gdextension"
	"graphics.gd/variant/Object"
)

type GoScript struct {
	ScriptExtension.Extension[GoScript] `gd:"GoScript"`
}

func (script *GoScript) EditorCanReloadFromFile() bool                     { return true }
func (script *GoScript) PlaceholderErased(placeholder gdextension.Pointer) {}
func (script *GoScript) CanInstantiate() bool                              { return false }
func (script *GoScript) GetBaseScript() Script.Instance                    { return script.AsScript() }
func (script *GoScript) GetGlobalName() string                             { return "gdscript" }
func (script *GoScript) InheritsScript(parent Script.Instance) bool        { return false }
func (script *GoScript) GetInstanceBaseType() string {
	return "Node"
}
func (script *GoScript) InstanceCreate(obj Object.Instance) gdextension.Pointer {
	return 0
}
func (script *GoScript) PlaceholderInstanceCreate(for_object Object.Instance) gdextension.Pointer {
	return 0
}
func (script *GoScript) InstanceHas(obj Object.Instance) bool {
	return false
}
func (script *GoScript) HasSourceCode() bool {
	return false
}
func (script *GoScript) GetSourceCode() string {
	return ""
}
func (script *GoScript) SetSourceCode(code string) {

}
func (script *GoScript) Reload(keep_state bool) error {
	return nil
}
func (script *GoScript) GetDocClassName() string {
	return ""
}
func (script *GoScript) GetDocumentation() [][]ScriptExtension.ClassDoc {
	return nil
}
func (script *GoScript) GetClassIconPath() string {
	return "res://icon.png"
}
func (script *GoScript) HasMethod(method string) bool {
	return false
}
func (script *GoScript) HasStaticMethod(method string) bool {
	return false
}

// Return the expected argument count for the given 'method', or null if it can't be determined (which will then fall back to the default behavior).
func (script *GoScript) GetScriptMethodArgumentCount(method string) any {
	return nil
}
func (script *GoScript) GetMethodInfo(method string) Object.MethodInfo {
	return Object.MethodInfo{}
}
func (script *GoScript) IsTool() bool {
	return false
}
func (script *GoScript) IsValid() bool {
	return true
}

// Returns true if the script is an abstract script. Abstract scripts cannot be instantiated directly, instead other scripts should inherit them. Abstract scripts will be either unselectable or hidden in the Create New Node dialog (unselectable if there are non-abstract classes inheriting it, otherwise hidden).
func (script *GoScript) IsAbstract() bool {
	return false
}
func (script *GoScript) GetLanguage() ScriptLanguage.Instance {
	return ScriptLanguage.Nil
}
func (script *GoScript) HasScriptSignal(signal string) bool {
	return false
}
func (script *GoScript) GetScriptSignalList() [][]struct{} {
	return nil
}
func (script *GoScript) HasPropertyDefaultValue(property string) bool {
	return false
}
func (script *GoScript) GetPropertyDefaultValue(property string) any {
	return nil
}
func (script *GoScript) UpdateExports() {}
func (script *GoScript) GetScriptMethodList() [][]struct{} {
	return nil
}
func (script *GoScript) GetScriptPropertyList() [][]struct{} {
	return nil
}
func (script *GoScript) GetMemberLine(member string) int {
	return 0
}
func (script *GoScript) GetConstants() map[string]any {
	return nil
}
func (script *GoScript) GetMembers() []string {
	return nil
}
func (script *GoScript) IsPlaceholderFallbackEnabled() bool {
	return false
}
func (script *GoScript) GetRpcConfig() any {
	return nil
}
