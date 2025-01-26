// Package EditorCommandPalette provides methods for working with EditorCommandPalette object instances.
package EditorCommandPalette

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/ConfirmationDialog"
import "graphics.gd/classdb/AcceptDialog"
import "graphics.gd/classdb/Window"
import "graphics.gd/classdb/Viewport"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any

/*
Object that holds all the available Commands and their shortcuts text. These Commands can be accessed through [b]Editor > Command Palette[/b] menu.
Command key names use slash delimiters to distinguish sections, for example: [code]"example/command1"[/code] then [code]example[/code] will be the section name.
[codeblocks]
[gdscript]
var command_palette = EditorInterface.get_command_palette()
# external_command is a function that will be called with the command is executed.
var command_callable = Callable(self, "external_command").bind(arguments)
command_palette.add_command("command", "test/command",command_callable)
[/gdscript]
[csharp]
EditorCommandPalette commandPalette = EditorInterface.Singleton.GetCommandPalette();
// ExternalCommand is a function that will be called with the command is executed.
Callable commandCallable = new Callable(this, MethodName.ExternalCommand);
commandPalette.AddCommand("command", "test/command", commandCallable)
[/csharp]
[/codeblocks]
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_command_palette].
*/
type Instance [1]gdclass.EditorCommandPalette

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorCommandPalette() Instance
}

/*
Adds a custom command to EditorCommandPalette.
- [param command_name]: [String] (Name of the [b]Command[/b]. This is displayed to the user.)
- [param key_name]: [String] (Name of the key for a particular [b]Command[/b]. This is used to uniquely identify the [b]Command[/b].)
- [param binded_callable]: [Callable] (Callable of the [b]Command[/b]. This will be executed when the [b]Command[/b] is selected.)
- [param shortcut_text]: [String] (Shortcut text of the [b]Command[/b] if available.)
*/
func (self Instance) AddCommand(command_name string, key_name string, binded_callable func()) { //gd:EditorCommandPalette.add_command
	class(self).AddCommand(gd.NewString(command_name), gd.NewString(key_name), Callable.New(binded_callable), gd.NewString("None"))
}

/*
Removes the custom command from EditorCommandPalette.
- [param key_name]: [String] (Name of the key for a particular [b]Command[/b].)
*/
func (self Instance) RemoveCommand(key_name string) { //gd:EditorCommandPalette.remove_command
	class(self).RemoveCommand(gd.NewString(key_name))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorCommandPalette

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorCommandPalette"))
	casted := Instance{*(*gdclass.EditorCommandPalette)(unsafe.Pointer(&object))}
	return casted
}

/*
Adds a custom command to EditorCommandPalette.
- [param command_name]: [String] (Name of the [b]Command[/b]. This is displayed to the user.)
- [param key_name]: [String] (Name of the key for a particular [b]Command[/b]. This is used to uniquely identify the [b]Command[/b].)
- [param binded_callable]: [Callable] (Callable of the [b]Command[/b]. This will be executed when the [b]Command[/b] is selected.)
- [param shortcut_text]: [String] (Shortcut text of the [b]Command[/b] if available.)
*/
//go:nosplit
func (self class) AddCommand(command_name gd.String, key_name gd.String, binded_callable Callable.Function, shortcut_text gd.String) { //gd:EditorCommandPalette.add_command
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(command_name))
	callframe.Arg(frame, pointers.Get(key_name))
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(binded_callable)))
	callframe.Arg(frame, pointers.Get(shortcut_text))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorCommandPalette.Bind_add_command, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the custom command from EditorCommandPalette.
- [param key_name]: [String] (Name of the key for a particular [b]Command[/b].)
*/
//go:nosplit
func (self class) RemoveCommand(key_name gd.String) { //gd:EditorCommandPalette.remove_command
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key_name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorCommandPalette.Bind_remove_command, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsEditorCommandPalette() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorCommandPalette() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsConfirmationDialog() ConfirmationDialog.Advanced {
	return *((*ConfirmationDialog.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsConfirmationDialog() ConfirmationDialog.Instance {
	return *((*ConfirmationDialog.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAcceptDialog() AcceptDialog.Advanced {
	return *((*AcceptDialog.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAcceptDialog() AcceptDialog.Instance {
	return *((*AcceptDialog.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWindow() Window.Advanced    { return *((*Window.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWindow() Window.Instance { return *((*Window.Instance)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.Advanced {
	return *((*Viewport.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsViewport() Viewport.Instance {
	return *((*Viewport.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ConfirmationDialog.Advanced(self.AsConfirmationDialog()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(ConfirmationDialog.Instance(self.AsConfirmationDialog()), name)
	}
}
func init() {
	gdclass.Register("EditorCommandPalette", func(ptr gd.Object) any {
		return [1]gdclass.EditorCommandPalette{*(*gdclass.EditorCommandPalette)(unsafe.Pointer(&ptr))}
	})
}
