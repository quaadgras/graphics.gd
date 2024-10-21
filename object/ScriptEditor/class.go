package ScriptEditor

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PanelContainer"
import "grow.graphics/gd/object/Container"
import "grow.graphics/gd/object/Control"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Godot editor's script editor.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_script_editor].

*/
type Simple [1]classdb.ScriptEditor
func (self Simple) GetCurrentEditor() [1]classdb.ScriptEditorBase {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.ScriptEditorBase(Expert(self).GetCurrentEditor(gc))
}
func (self Simple) GetOpenScriptEditors() gd.ArrayOf[[1]classdb.ScriptEditorBase] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.ScriptEditorBase](Expert(self).GetOpenScriptEditors(gc))
}
func (self Simple) RegisterSyntaxHighlighter(syntax_highlighter [1]classdb.EditorSyntaxHighlighter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterSyntaxHighlighter(syntax_highlighter)
}
func (self Simple) UnregisterSyntaxHighlighter(syntax_highlighter [1]classdb.EditorSyntaxHighlighter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnregisterSyntaxHighlighter(syntax_highlighter)
}
func (self Simple) GotoLine(line_number int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GotoLine(gd.Int(line_number))
}
func (self Simple) GetCurrentScript() [1]classdb.Script {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Script(Expert(self).GetCurrentScript(gc))
}
func (self Simple) GetOpenScripts() gd.ArrayOf[[1]classdb.Script] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[[1]classdb.Script](Expert(self).GetOpenScripts(gc))
}
func (self Simple) OpenScriptCreateDialog(base_name string, base_path string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).OpenScriptCreateDialog(gc.String(base_name), gc.String(base_path))
}
func (self Simple) GotoHelp(topic string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GotoHelp(gc.String(topic))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ScriptEditor
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the [ScriptEditorBase] object that the user is currently editing.
*/
//go:nosplit
func (self class) GetCurrentEditor(ctx gd.Lifetime) object.ScriptEditorBase {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_current_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.ScriptEditorBase
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an array with all [ScriptEditorBase] objects which are currently open in editor.
*/
//go:nosplit
func (self class) GetOpenScriptEditors(ctx gd.Lifetime) gd.ArrayOf[object.ScriptEditorBase] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_open_script_editors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.ScriptEditorBase](ret)
}
/*
Registers the [EditorSyntaxHighlighter] to the editor, the [EditorSyntaxHighlighter] will be available on all open scripts.
[b]Note:[/b] Does not apply to scripts that are already opened.
*/
//go:nosplit
func (self class) RegisterSyntaxHighlighter(syntax_highlighter object.EditorSyntaxHighlighter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(syntax_highlighter[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_register_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters the [EditorSyntaxHighlighter] from the editor.
[b]Note:[/b] The [EditorSyntaxHighlighter] will still be applied to scripts that are already opened.
*/
//go:nosplit
func (self class) UnregisterSyntaxHighlighter(syntax_highlighter object.EditorSyntaxHighlighter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(syntax_highlighter[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_unregister_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Goes to the specified line in the current script.
*/
//go:nosplit
func (self class) GotoLine(line_number gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line_number)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_goto_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Script] that is currently active in editor.
*/
//go:nosplit
func (self class) GetCurrentScript(ctx gd.Lifetime) object.Script {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_current_script, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Script
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an array with all [Script] objects which are currently open in editor.
*/
//go:nosplit
func (self class) GetOpenScripts(ctx gd.Lifetime) gd.ArrayOf[object.Script] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_open_scripts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[object.Script](ret)
}
/*
Opens the script create dialog. The script will extend [param base_name]. The file extension can be omitted from [param base_path]. It will be added based on the selected scripting language.
*/
//go:nosplit
func (self class) OpenScriptCreateDialog(base_name gd.String, base_path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base_name))
	callframe.Arg(frame, mmm.Get(base_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_open_script_create_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Opens help for the given topic. The [param topic] is an encoded string that controls which class, method, constant, signal, annotation, property, or theme item should be focused.
The supported [param topic] formats include [code]class_name:class[/code], [code]class_method:class:method[/code], [code]class_constant:class:constant[/code], [code]class_signal:class:signal[/code], [code]class_annotation:class:@annotation[/code], [code]class_property:class:property[/code], and [code]class_theme_item:class:item[/code], where [code]class[/code] is the class name, [code]method[/code] is the method name, [code]constant[/code] is the constant name, [code]signal[/code] is the signal name, [code]annotation[/code] is the annotation name, [code]property[/code] is the property name, and [code]item[/code] is the theme item.
[b]Examples:[/b]
[codeblock]
# Shows help for the Node class.
class_name:Node
# Shows help for the global min function.
# Global objects are accessible in the `@GlobalScope` namespace, shown here.
class_method:@GlobalScope:min
# Shows help for get_viewport in the Node class.
class_method:Node:get_viewport
# Shows help for the Input constant MOUSE_BUTTON_MIDDLE.
class_constant:Input:MOUSE_BUTTON_MIDDLE
# Shows help for the BaseButton signal pressed.
class_signal:BaseButton:pressed
# Shows help for the CanvasItem property visible.
class_property:CanvasItem:visible
# Shows help for the GDScript annotation export.
# Annotations should be prefixed with the `@` symbol in the descriptor, as shown here.
class_annotation:@GDScript:@export
# Shows help for the GraphNode theme item named panel_selected.
class_theme_item:GraphNode:panel_selected
[/codeblock]
*/
//go:nosplit
func (self class) GotoHelp(topic gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(topic))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_goto_help, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsScriptEditor() Expert { return self[0].AsScriptEditor() }


//go:nosplit
func (self Simple) AsScriptEditor() Simple { return self[0].AsScriptEditor() }


//go:nosplit
func (self class) AsPanelContainer() PanelContainer.Expert { return self[0].AsPanelContainer() }


//go:nosplit
func (self Simple) AsPanelContainer() PanelContainer.Simple { return self[0].AsPanelContainer() }


//go:nosplit
func (self class) AsContainer() Container.Expert { return self[0].AsContainer() }


//go:nosplit
func (self Simple) AsContainer() Container.Simple { return self[0].AsContainer() }


//go:nosplit
func (self class) AsControl() Control.Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Control.Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ScriptEditor", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
