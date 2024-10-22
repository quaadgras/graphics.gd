package ScriptEditorBase

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
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
var _ mmm.Lifetime

/*
Base editor for editing scripts in the [ScriptEditor]. This does not include documentation items.

*/
type Go [1]classdb.ScriptEditorBase

/*
Returns the underlying [Control] used for editing scripts. For text scripts, this is a [CodeEdit].
*/
func (self Go) GetBaseEditor() gdclass.Control {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Control(class(self).GetBaseEditor(gc))
}

/*
Adds a [EditorSyntaxHighlighter] to the open script.
*/
func (self Go) AddSyntaxHighlighter(highlighter gdclass.EditorSyntaxHighlighter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddSyntaxHighlighter(highlighter)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ScriptEditorBase
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ScriptEditorBase"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the underlying [Control] used for editing scripts. For text scripts, this is a [CodeEdit].
*/
//go:nosplit
func (self class) GetBaseEditor(ctx gd.Lifetime) gdclass.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditorBase.Bind_get_base_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Adds a [EditorSyntaxHighlighter] to the open script.
*/
//go:nosplit
func (self class) AddSyntaxHighlighter(highlighter gdclass.EditorSyntaxHighlighter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(highlighter[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditorBase.Bind_add_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnNameChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("name_changed"), gc.Callable(cb), 0)
}


func (self Go) OnEditedScriptChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("edited_script_changed"), gc.Callable(cb), 0)
}


func (self Go) OnRequestHelp(cb func(topic string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("request_help"), gc.Callable(cb), 0)
}


func (self Go) OnRequestOpenScriptAtLine(cb func(script gd.Object, line int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("request_open_script_at_line"), gc.Callable(cb), 0)
}


func (self Go) OnRequestSaveHistory(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("request_save_history"), gc.Callable(cb), 0)
}


func (self Go) OnRequestSavePreviousState(cb func(state gd.Dictionary)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("request_save_previous_state"), gc.Callable(cb), 0)
}


func (self Go) OnGoToHelp(cb func(what string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("go_to_help"), gc.Callable(cb), 0)
}


func (self Go) OnSearchInFilesRequested(cb func(text string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("search_in_files_requested"), gc.Callable(cb), 0)
}


func (self Go) OnReplaceInFilesRequested(cb func(text string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("replace_in_files_requested"), gc.Callable(cb), 0)
}


func (self Go) OnGoToMethod(cb func(script gd.Object, method string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("go_to_method"), gc.Callable(cb), 0)
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
func init() {classdb.Register("ScriptEditorBase", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
