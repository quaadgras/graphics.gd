// Package EditorDebuggerPlugin provides methods for working with EditorDebuggerPlugin object instances.
package EditorDebuggerPlugin

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[EditorDebuggerPlugin] provides functions related to the editor side of the debugger.
To interact with the debugger, an instance of this class must be added to the editor via [method EditorPlugin.add_debugger_plugin].
Once added, the [method _setup_session] callback will be called for every [EditorDebuggerSession] available to the plugin, and when new ones are created (the sessions may be inactive during this stage).
You can retrieve the available [EditorDebuggerSession]s via [method get_sessions] or get a specific one via [method get_session].
[codeblocks]
[gdscript]
@tool
extends EditorPlugin

class ExampleEditorDebugger extends EditorDebuggerPlugin:

	func _has_capture(capture):
	    # Return true if you wish to handle messages with the prefix "my_plugin:".
	    return capture == "my_plugin"

	func _capture(message, data, session_id):
	    if message == "my_plugin:ping":
	        get_session(session_id).send_message("my_plugin:echo", data)
	        return true
	    return false

	func _setup_session(session_id):
	    # Add a new tab in the debugger session UI containing a label.
	    var label = Label.new()
	    label.name = "Example plugin" # Will be used as the tab title.
	    label.text = "Example plugin"
	    var session = get_session(session_id)
	    # Listens to the session started and stopped signals.
	    session.started.connect(func (): print("Session started"))
	    session.stopped.connect(func (): print("Session stopped"))
	    session.add_session_tab(label)

var debugger = ExampleEditorDebugger.new()

func _enter_tree():

	add_debugger_plugin(debugger)

func _exit_tree():

	remove_debugger_plugin(debugger)

[/gdscript]
[/codeblocks]
To connect on the running game side, use the [EngineDebugger] singleton:
[codeblocks]
[gdscript]
extends Node

func _ready():

	EngineDebugger.register_message_capture("my_plugin", _capture)
	EngineDebugger.send_message("my_plugin:ping", ["test"])

func _capture(message, data):

	# Note that the "my_plugin:" prefix is not used here.
	if message == "echo":
	    prints("Echo received:", data)
	    return true
	return false

[/gdscript]
[/codeblocks]
[b]Note:[/b] While the game is running, [method @GlobalScope.print] and similar functions [i]called in the editor[/i] do not print anything, the Output Log prints only game messages.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorDebuggerPlugin)
*/
type Instance [1]gdclass.EditorDebuggerPlugin

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorDebuggerPlugin() Instance
}
type Interface interface {
	//Override this method to be notified whenever a new [EditorDebuggerSession] is created. Note that the session may be inactive during this stage.
	SetupSession(session_id int)
	//Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passed to the [method _capture] method.
	HasCapture(capture string) bool
	//Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the [param message]. Use [method get_session] to retrieve the session. This method should return [code]true[/code] if the message is recognized.
	Capture(message string, data []any, session_id int) bool
	//Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
	GotoScriptLine(script [1]gdclass.Script, line int)
	//Override this method to be notified when all breakpoints are cleared in the editor.
	BreakpointsClearedInTree()
	//Override this method to be notified when a breakpoint is set in the editor.
	BreakpointSetInTree(script [1]gdclass.Script, line int, enabled bool)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) SetupSession(session_id int)                                 { return }
func (self implementation) HasCapture(capture string) (_ bool)                          { return }
func (self implementation) Capture(message string, data []any, session_id int) (_ bool) { return }
func (self implementation) GotoScriptLine(script [1]gdclass.Script, line int)           { return }
func (self implementation) BreakpointsClearedInTree()                                   { return }
func (self implementation) BreakpointSetInTree(script [1]gdclass.Script, line int, enabled bool) {
	return
}

/*
Override this method to be notified whenever a new [EditorDebuggerSession] is created. Note that the session may be inactive during this stage.
*/
func (Instance) _setup_session(impl func(ptr unsafe.Pointer, session_id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var session_id = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(session_id))
	}
}

/*
Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passed to the [method _capture] method.
*/
func (Instance) _has_capture(impl func(ptr unsafe.Pointer, capture string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var capture = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(capture))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, capture.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the [param message]. Use [method get_session] to retrieve the session. This method should return [code]true[/code] if the message is recognized.
*/
func (Instance) _capture(impl func(ptr unsafe.Pointer, message string, data []any, session_id int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var message = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(message))
		var data = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalArray(data))
		var session_id = gd.UnsafeGet[int64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, message.String(), gd.ArrayAs[[]any](gd.InternalArray(data)), int(session_id))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
*/
func (Instance) _goto_script_line(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, line int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var line = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, int(line))
	}
}

/*
Override this method to be notified when all breakpoints are cleared in the editor.
*/
func (Instance) _breakpoints_cleared_in_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to be notified when a breakpoint is set in the editor.
*/
func (Instance) _breakpoint_set_in_tree(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, line int, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var line = gd.UnsafeGet[int64](p_args, 1)

		var enabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, int(line), enabled)
	}
}

/*
Returns the [EditorDebuggerSession] with the given [param id].
*/
func (self Instance) GetSession(id int) [1]gdclass.EditorDebuggerSession { //gd:EditorDebuggerPlugin.get_session
	return [1]gdclass.EditorDebuggerSession(Advanced(self).GetSession(int64(id)))
}

/*
Returns an array of [EditorDebuggerSession] currently available to this debugger plugin.
[b]Note:[/b] Sessions in the array may be inactive, check their state via [method EditorDebuggerSession.is_active].
*/
func (self Instance) GetSessions() []any { //gd:EditorDebuggerPlugin.get_sessions
	return []any(gd.ArrayAs[[]any](gd.InternalArray(Advanced(self).GetSessions())))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorDebuggerPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorDebuggerPlugin"))
	casted := Instance{*(*gdclass.EditorDebuggerPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override this method to be notified whenever a new [EditorDebuggerSession] is created. Note that the session may be inactive during this stage.
*/
func (class) _setup_session(impl func(ptr unsafe.Pointer, session_id int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var session_id = gd.UnsafeGet[int64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, session_id)
	}
}

/*
Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passed to the [method _capture] method.
*/
func (class) _has_capture(impl func(ptr unsafe.Pointer, capture String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var capture = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(capture))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, capture)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the [param message]. Use [method get_session] to retrieve the session. This method should return [code]true[/code] if the message is recognized.
*/
func (class) _capture(impl func(ptr unsafe.Pointer, message String.Readable, data Array.Any, session_id int64) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var message = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(message))
		var data = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalArray(data))
		var session_id = gd.UnsafeGet[int64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, message, data, session_id)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
*/
func (class) _goto_script_line(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, line int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var line = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, line)
	}
}

/*
Override this method to be notified when all breakpoints are cleared in the editor.
*/
func (class) _breakpoints_cleared_in_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to be notified when a breakpoint is set in the editor.
*/
func (class) _breakpoint_set_in_tree(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, line int64, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var line = gd.UnsafeGet[int64](p_args, 1)

		var enabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, line, enabled)
	}
}

/*
Returns the [EditorDebuggerSession] with the given [param id].
*/
//go:nosplit
func (self class) GetSession(id int64) [1]gdclass.EditorDebuggerSession { //gd:EditorDebuggerPlugin.get_session
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerPlugin.Bind_get_session, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorDebuggerSession{gd.PointerWithOwnershipTransferredToGo[gdclass.EditorDebuggerSession](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an array of [EditorDebuggerSession] currently available to this debugger plugin.
[b]Note:[/b] Sessions in the array may be inactive, check their state via [method EditorDebuggerSession.is_active].
*/
//go:nosplit
func (self class) GetSessions() Array.Any { //gd:EditorDebuggerPlugin.get_sessions
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerPlugin.Bind_get_sessions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsEditorDebuggerPlugin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorDebuggerPlugin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_session":
		return reflect.ValueOf(self._setup_session)
	case "_has_capture":
		return reflect.ValueOf(self._has_capture)
	case "_capture":
		return reflect.ValueOf(self._capture)
	case "_goto_script_line":
		return reflect.ValueOf(self._goto_script_line)
	case "_breakpoints_cleared_in_tree":
		return reflect.ValueOf(self._breakpoints_cleared_in_tree)
	case "_breakpoint_set_in_tree":
		return reflect.ValueOf(self._breakpoint_set_in_tree)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_session":
		return reflect.ValueOf(self._setup_session)
	case "_has_capture":
		return reflect.ValueOf(self._has_capture)
	case "_capture":
		return reflect.ValueOf(self._capture)
	case "_goto_script_line":
		return reflect.ValueOf(self._goto_script_line)
	case "_breakpoints_cleared_in_tree":
		return reflect.ValueOf(self._breakpoints_cleared_in_tree)
	case "_breakpoint_set_in_tree":
		return reflect.ValueOf(self._breakpoint_set_in_tree)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorDebuggerPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorDebuggerPlugin{*(*gdclass.EditorDebuggerPlugin)(unsafe.Pointer(&ptr))}
	})
}
