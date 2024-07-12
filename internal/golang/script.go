package golang

import (
	"unsafe"

	"grow.graphics/gd"
)

type Script struct {
	gd.Class[Script, gd.ScriptExtension] `gd:"GoScript"`
}

func (script *Script) CanInstantiate() gd.Bool {
	return false
}

func (script *Script) EditorCanReloadFromFile() gd.Bool {
	return true
}

func (script *Script) GetBaseScript() gd.Script {
	return script.Super().AsScript()
}

func (script *Script) GetClassIconPath() gd.String {
	return script.Temporary.String("res://icon.png")
}

func (script *Script) GetConstants() gd.Dictionary {
	return script.Temporary.Dictionary()
}

func (script *Script) GetDocumentation() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *Script) GetGlobalName() gd.StringName {
	return script.Temporary.StringName("gdscript")
}

func (script *Script) GetInstanceBaseType() gd.StringName {
	return script.Temporary.StringName("Node")
}

func (script *Script) GetLanguage() gd.ScriptLanguage {
	return gd.ScriptLanguage{}
}

func (script *Script) GetMemberLine(member gd.StringName) gd.Int {
	return 0
}

func (script *Script) GetMembers() gd.ArrayOf[gd.StringName] {
	return nil
}

func (script *Script) GetMethodInfo() gd.Dictionary {
	return script.Temporary.Dictionary()
}

func (script *Script) GetPropertyDefaultValue(property gd.StringName) gd.Variant {
	return gd.Variant{}
}

func (script *Script) GetRpcConfig() gd.Variant {
	return gd.Variant{}
}

func (script *Script) GetScriptMethodList() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *Script) GetScriptPropertyList() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *Script) GetScriptSignalList() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *Script) GetSourceCode() gd.String {
	return script.Temporary.String("")
}

func (script *Script) HasMethod(method gd.StringName) gd.Bool {
	return false
}

func (script *Script) HasPropertyDefaultValue(property gd.StringName) gd.Bool {
	return false
}

func (script *Script) HasScriptSignal(signal gd.StringName) gd.Bool {
	return false
}

func (script *Script) HasSourceCode() gd.Bool {
	return false
}

func (script *Script) HasStaticMethod(method gd.StringName) gd.Bool {
	return false
}

func (script *Script) InheritsScript(parent Script) gd.Bool {
	return false
}

func (script *Script) InstanceCreate(obj gd.Object) unsafe.Pointer {
	return nil
}

func (script *Script) InstanceHas(obj gd.Object) gd.Bool {
	return false
}

func (script *Script) IsAbstract() gd.Bool {
	return false
}

func (script *Script) IsPlaceholderFallbackEnabled() gd.Bool {
	return false
}

func (script *Script) IsTool() gd.Bool {
	return false
}

func (script *Script) IsValid() gd.Bool {
	return true
}

func (script *Script) PlaceholderErased(placeholder unsafe.Pointer) unsafe.Pointer {
	return nil
}

func (script *Script) PlaceholderInstanceCraete(obj gd.Object) unsafe.Pointer {
	return nil
}

func (script *Script) Reload(keep_state bool) gd.Int {
	return 0
}

func (script *Script) SetSourceCode(code gd.String) {

}

func (script *Script) UpdateExports() {}
