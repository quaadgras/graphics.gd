package golang

import (
	"unsafe"

	"grow.graphics/gd"

	"grow.graphics/gd/object/Script"
	"grow.graphics/gd/object/ScriptExtension"
	"grow.graphics/gd/object/ScriptLanguage"
)

type GoScript struct {
	gd.Class[GoScript, ScriptExtension.Expert] `gd:"GoScript"`
}

func (script *GoScript) CanInstantiate() gd.Bool {
	return false
}

func (script *GoScript) EditorCanReloadFromFile() gd.Bool {
	return true
}

func (script *GoScript) GetBaseScript() Script.Expert {
	return script.Super().AsScript()
}

func (script *GoScript) GetClassIconPath() gd.String {
	return script.Temporary.String("res://icon.png")
}

func (script *GoScript) GetConstants() gd.Dictionary {
	return script.Temporary.Dictionary()
}

func (script *GoScript) GetDocumentation() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *GoScript) GetGlobalName() gd.StringName {
	return script.Temporary.StringName("gdscript")
}

func (script *GoScript) GetInstanceBaseType() gd.StringName {
	return script.Temporary.StringName("Node")
}

func (script *GoScript) GetLanguage() ScriptLanguage.Expert {
	return ScriptLanguage.Expert{}
}

func (script *GoScript) GetMemberLine(member gd.StringName) gd.Int {
	return 0
}

func (script *GoScript) GetMembers() gd.ArrayOf[gd.StringName] {
	return nil
}

func (script *GoScript) GetMethodInfo() gd.Dictionary {
	return script.Temporary.Dictionary()
}

func (script *GoScript) GetPropertyDefaultValue(property gd.StringName) gd.Variant {
	return gd.Variant{}
}

func (script *GoScript) GetRpcConfig() gd.Variant {
	return gd.Variant{}
}

func (script *GoScript) GetScriptMethodList() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *GoScript) GetScriptPropertyList() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *GoScript) GetScriptSignalList() gd.ArrayOf[gd.Dictionary] {
	return nil
}

func (script *GoScript) GetSourceCode() gd.String {
	return script.Temporary.String("")
}

func (script *GoScript) HasMethod(method gd.StringName) gd.Bool {
	return false
}

func (script *GoScript) HasPropertyDefaultValue(property gd.StringName) gd.Bool {
	return false
}

func (script *GoScript) HasScriptSignal(signal gd.StringName) gd.Bool {
	return false
}

func (script *GoScript) HasSourceCode() gd.Bool {
	return false
}

func (script *GoScript) HasStaticMethod(method gd.StringName) gd.Bool {
	return false
}

func (script *GoScript) InheritsScript(parent GoScript) gd.Bool {
	return false
}

func (script *GoScript) InstanceCreate(obj gd.Object) unsafe.Pointer {
	return nil
}

func (script *GoScript) InstanceHas(obj gd.Object) gd.Bool {
	return false
}

func (script *GoScript) IsAbstract() gd.Bool {
	return false
}

func (script *GoScript) IsPlaceholderFallbackEnabled() gd.Bool {
	return false
}

func (script *GoScript) IsTool() gd.Bool {
	return false
}

func (script *GoScript) IsValid() gd.Bool {
	return true
}

func (script *GoScript) PlaceholderErased(placeholder unsafe.Pointer) unsafe.Pointer {
	return nil
}

func (script *GoScript) PlaceholderInstanceCraete(obj gd.Object) unsafe.Pointer {
	return nil
}

func (script *GoScript) Reload(keep_state bool) gd.Int {
	return 0
}

func (script *GoScript) SetSourceCode(code gd.String) {

}

func (script *GoScript) UpdateExports() {}
