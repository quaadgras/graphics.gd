package ScriptEditorBase

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VBoxContainer"
import "grow.graphics/gd/gdclass/BoxContainer"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Base editor for editing scripts in the [ScriptEditor]. This does not include documentation items.

*/
type Go [1]classdb.ScriptEditorBase

/*
Returns the underlying [Control] used for editing scripts. For text scripts, this is a [CodeEdit].
*/
func (self Go) GetBaseEditor() gdclass.Control {
	return gdclass.Control(class(self).GetBaseEditor())
}

/*
Adds a [EditorSyntaxHighlighter] to the open script.
*/
func (self Go) AddSyntaxHighlighter(highlighter gdclass.EditorSyntaxHighlighter) {
	class(self).AddSyntaxHighlighter(highlighter)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ScriptEditorBase
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptEditorBase"))
	return Go{classdb.ScriptEditorBase(object)}
}

/*
Returns the underlying [Control] used for editing scripts. For text scripts, this is a [CodeEdit].
*/
//go:nosplit
func (self class) GetBaseEditor() gdclass.Control {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditorBase.Bind_get_base_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Control{classdb.Control(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Adds a [EditorSyntaxHighlighter] to the open script.
*/
//go:nosplit
func (self class) AddSyntaxHighlighter(highlighter gdclass.EditorSyntaxHighlighter)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(highlighter[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditorBase.Bind_add_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnNameChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("name_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnEditedScriptChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("edited_script_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnRequestHelp(cb func(topic string)) {
	self[0].AsObject().Connect(gd.NewStringName("request_help"), gd.NewCallable(cb), 0)
}


func (self Go) OnRequestOpenScriptAtLine(cb func(script gd.Object, line int)) {
	self[0].AsObject().Connect(gd.NewStringName("request_open_script_at_line"), gd.NewCallable(cb), 0)
}


func (self Go) OnRequestSaveHistory(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("request_save_history"), gd.NewCallable(cb), 0)
}


func (self Go) OnRequestSavePreviousState(cb func(state gd.Dictionary)) {
	self[0].AsObject().Connect(gd.NewStringName("request_save_previous_state"), gd.NewCallable(cb), 0)
}


func (self Go) OnGoToHelp(cb func(what string)) {
	self[0].AsObject().Connect(gd.NewStringName("go_to_help"), gd.NewCallable(cb), 0)
}


func (self Go) OnSearchInFilesRequested(cb func(text string)) {
	self[0].AsObject().Connect(gd.NewStringName("search_in_files_requested"), gd.NewCallable(cb), 0)
}


func (self Go) OnReplaceInFilesRequested(cb func(text string)) {
	self[0].AsObject().Connect(gd.NewStringName("replace_in_files_requested"), gd.NewCallable(cb), 0)
}


func (self Go) OnGoToMethod(cb func(script gd.Object, method string)) {
	self[0].AsObject().Connect(gd.NewStringName("go_to_method"), gd.NewCallable(cb), 0)
}


func (self class) AsScriptEditorBase() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsScriptEditorBase() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVBoxContainer() VBoxContainer.GD { return *((*VBoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVBoxContainer() VBoxContainer.Go { return *((*VBoxContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsBoxContainer() BoxContainer.GD { return *((*BoxContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsBoxContainer() BoxContainer.Go { return *((*BoxContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVBoxContainer(), name)
	}
}
func init() {classdb.Register("ScriptEditorBase", func(ptr gd.Object) any { return classdb.ScriptEditorBase(ptr) })}
