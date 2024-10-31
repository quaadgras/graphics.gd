package ScriptEditor

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PanelContainer"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Godot editor's script editor.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_script_editor].
*/
type Instance [1]classdb.ScriptEditor

/*
Returns the [ScriptEditorBase] object that the user is currently editing.
*/
func (self Instance) GetCurrentEditor() gdclass.ScriptEditorBase {
	return gdclass.ScriptEditorBase(class(self).GetCurrentEditor())
}

/*
Returns an array with all [ScriptEditorBase] objects which are currently open in editor.
*/
func (self Instance) GetOpenScriptEditors() gd.Array {
	return gd.Array(class(self).GetOpenScriptEditors())
}

/*
Registers the [EditorSyntaxHighlighter] to the editor, the [EditorSyntaxHighlighter] will be available on all open scripts.
[b]Note:[/b] Does not apply to scripts that are already opened.
*/
func (self Instance) RegisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter) {
	class(self).RegisterSyntaxHighlighter(syntax_highlighter)
}

/*
Unregisters the [EditorSyntaxHighlighter] from the editor.
[b]Note:[/b] The [EditorSyntaxHighlighter] will still be applied to scripts that are already opened.
*/
func (self Instance) UnregisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter) {
	class(self).UnregisterSyntaxHighlighter(syntax_highlighter)
}

/*
Goes to the specified line in the current script.
*/
func (self Instance) GotoLine(line_number int) {
	class(self).GotoLine(gd.Int(line_number))
}

/*
Returns a [Script] that is currently active in editor.
*/
func (self Instance) GetCurrentScript() gdclass.Script {
	return gdclass.Script(class(self).GetCurrentScript())
}

/*
Returns an array with all [Script] objects which are currently open in editor.
*/
func (self Instance) GetOpenScripts() gd.Array {
	return gd.Array(class(self).GetOpenScripts())
}

/*
Opens the script create dialog. The script will extend [param base_name]. The file extension can be omitted from [param base_path]. It will be added based on the selected scripting language.
*/
func (self Instance) OpenScriptCreateDialog(base_name string, base_path string) {
	class(self).OpenScriptCreateDialog(gd.NewString(base_name), gd.NewString(base_path))
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
func (self Instance) GotoHelp(topic string) {
	class(self).GotoHelp(gd.NewString(topic))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ScriptEditor

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptEditor"))
	return Instance{classdb.ScriptEditor(object)}
}

/*
Returns the [ScriptEditorBase] object that the user is currently editing.
*/
//go:nosplit
func (self class) GetCurrentEditor() gdclass.ScriptEditorBase {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_get_current_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.ScriptEditorBase{classdb.ScriptEditorBase(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns an array with all [ScriptEditorBase] objects which are currently open in editor.
*/
//go:nosplit
func (self class) GetOpenScriptEditors() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_get_open_script_editors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Registers the [EditorSyntaxHighlighter] to the editor, the [EditorSyntaxHighlighter] will be available on all open scripts.
[b]Note:[/b] Does not apply to scripts that are already opened.
*/
//go:nosplit
func (self class) RegisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(syntax_highlighter[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_register_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Unregisters the [EditorSyntaxHighlighter] from the editor.
[b]Note:[/b] The [EditorSyntaxHighlighter] will still be applied to scripts that are already opened.
*/
//go:nosplit
func (self class) UnregisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(syntax_highlighter[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_unregister_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Goes to the specified line in the current script.
*/
//go:nosplit
func (self class) GotoLine(line_number gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, line_number)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_goto_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [Script] that is currently active in editor.
*/
//go:nosplit
func (self class) GetCurrentScript() gdclass.Script {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_get_current_script, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Script{classdb.Script(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns an array with all [Script] objects which are currently open in editor.
*/
//go:nosplit
func (self class) GetOpenScripts() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_get_open_scripts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Opens the script create dialog. The script will extend [param base_name]. The file extension can be omitted from [param base_path]. It will be added based on the selected scripting language.
*/
//go:nosplit
func (self class) OpenScriptCreateDialog(base_name gd.String, base_path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base_name))
	callframe.Arg(frame, pointers.Get(base_path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_open_script_create_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) GotoHelp(topic gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(topic))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ScriptEditor.Bind_goto_help, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnEditorScriptChanged(cb func(script gdclass.Script)) {
	self[0].AsObject().Connect(gd.NewStringName("editor_script_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScriptClose(cb func(script gdclass.Script)) {
	self[0].AsObject().Connect(gd.NewStringName("script_close"), gd.NewCallable(cb), 0)
}

func (self class) AsScriptEditor() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScriptEditor() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPanelContainer() PanelContainer.Advanced {
	return *((*PanelContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPanelContainer() PanelContainer.Instance {
	return *((*PanelContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPanelContainer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPanelContainer(), name)
	}
}
func init() {
	classdb.Register("ScriptEditor", func(ptr gd.Object) any { return classdb.ScriptEditor(ptr) })
}
