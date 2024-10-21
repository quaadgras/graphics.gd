package EditorCommandPalette

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ConfirmationDialog"
import "grow.graphics/gd/object/AcceptDialog"
import "grow.graphics/gd/object/Window"
import "grow.graphics/gd/object/Viewport"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
type Simple [1]classdb.EditorCommandPalette
func (self Simple) AddCommand(command_name string, key_name string, binded_callable gd.Callable, shortcut_text string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCommand(gc.String(command_name), gc.String(key_name), binded_callable, gc.String(shortcut_text))
}
func (self Simple) RemoveCommand(key_name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveCommand(gc.String(key_name))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorCommandPalette
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Adds a custom command to EditorCommandPalette.
- [param command_name]: [String] (Name of the [b]Command[/b]. This is displayed to the user.)
- [param key_name]: [String] (Name of the key for a particular [b]Command[/b]. This is used to uniquely identify the [b]Command[/b].)
- [param binded_callable]: [Callable] (Callable of the [b]Command[/b]. This will be executed when the [b]Command[/b] is selected.)
- [param shortcut_text]: [String] (Shortcut text of the [b]Command[/b] if available.)
*/
//go:nosplit
func (self class) AddCommand(command_name gd.String, key_name gd.String, binded_callable gd.Callable, shortcut_text gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(command_name))
	callframe.Arg(frame, mmm.Get(key_name))
	callframe.Arg(frame, mmm.Get(binded_callable))
	callframe.Arg(frame, mmm.Get(shortcut_text))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorCommandPalette.Bind_add_command, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the custom command from EditorCommandPalette.
- [param key_name]: [String] (Name of the key for a particular [b]Command[/b].)
*/
//go:nosplit
func (self class) RemoveCommand(key_name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(key_name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorCommandPalette.Bind_remove_command, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorCommandPalette() Expert { return self[0].AsEditorCommandPalette() }


//go:nosplit
func (self Simple) AsEditorCommandPalette() Simple { return self[0].AsEditorCommandPalette() }


//go:nosplit
func (self class) AsConfirmationDialog() ConfirmationDialog.Expert { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self Simple) AsConfirmationDialog() ConfirmationDialog.Simple { return self[0].AsConfirmationDialog() }


//go:nosplit
func (self class) AsAcceptDialog() AcceptDialog.Expert { return self[0].AsAcceptDialog() }


//go:nosplit
func (self Simple) AsAcceptDialog() AcceptDialog.Simple { return self[0].AsAcceptDialog() }


//go:nosplit
func (self class) AsWindow() Window.Expert { return self[0].AsWindow() }


//go:nosplit
func (self Simple) AsWindow() Window.Simple { return self[0].AsWindow() }


//go:nosplit
func (self class) AsViewport() Viewport.Expert { return self[0].AsViewport() }


//go:nosplit
func (self Simple) AsViewport() Viewport.Simple { return self[0].AsViewport() }


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
func init() {classdb.Register("EditorCommandPalette", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
