package EditorDebuggerPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

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

	func _has_capture(prefix):
	    # Return true if you wish to handle message with this prefix.
	    return prefix == "my_plugin"

	func _capture(message, data, session_id):
	    if message == "my_plugin:ping":
	        get_session(session_id).send_message("my_plugin:echo", data)

	func _setup_session(session_id):
	    # Add a new tab in the debugger session UI containing a label.
	    var label = Label.new()
	    label.name = "Example plugin"
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

	// EditorDebuggerPlugin methods that can be overridden by a [Class] that extends it.
	type EditorDebuggerPlugin interface {
		//Override this method to be notified whenever a new [EditorDebuggerSession] is created (the session may be inactive during this stage).
		SetupSession(session_id int)
		//Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passes to the [method _capture] method.
		HasCapture(capture string) bool
		//Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the message (which you can retrieve via [method get_session]).
		Capture(message string, data gd.Array, session_id int) bool
		//Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
		GotoScriptLine(script objects.Script, line int)
		//Override this method to be notified when all breakpoints are cleared in the editor.
		BreakpointsClearedInTree()
		//Override this method to be notified when a breakpoint is set in the editor.
		BreakpointSetInTree(script objects.Script, line int, enabled bool)
	}
*/
type Instance [1]classdb.EditorDebuggerPlugin

/*
Override this method to be notified whenever a new [EditorDebuggerSession] is created (the session may be inactive during this stage).
*/
func (Instance) _setup_session(impl func(ptr unsafe.Pointer, session_id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var session_id = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(session_id))
	}
}

/*
Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passes to the [method _capture] method.
*/
func (Instance) _has_capture(impl func(ptr unsafe.Pointer, capture string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var capture = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(capture)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, capture.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the message (which you can retrieve via [method get_session]).
*/
func (Instance) _capture(impl func(ptr unsafe.Pointer, message string, data gd.Array, session_id int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var message = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(message)
		var data = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(data)
		var session_id = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, message.String(), data, int(session_id))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
*/
func (Instance) _goto_script_line(impl func(ptr unsafe.Pointer, script objects.Script, line int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var script = objects.Script{pointers.New[classdb.Script]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(script[0])
		var line = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, int(line))
	}
}

/*
Override this method to be notified when all breakpoints are cleared in the editor.
*/
func (Instance) _breakpoints_cleared_in_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to be notified when a breakpoint is set in the editor.
*/
func (Instance) _breakpoint_set_in_tree(impl func(ptr unsafe.Pointer, script objects.Script, line int, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var script = objects.Script{pointers.New[classdb.Script]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(script[0])
		var line = gd.UnsafeGet[gd.Int](p_args, 1)
		var enabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, int(line), enabled)
	}
}

/*
Returns the [EditorDebuggerSession] with the given [param id].
*/
func (self Instance) GetSession(id int) objects.EditorDebuggerSession {
	return objects.EditorDebuggerSession(class(self).GetSession(gd.Int(id)))
}

/*
Returns an array of [EditorDebuggerSession] currently available to this debugger plugin.
[b]Note:[/b] Sessions in the array may be inactive, check their state via [method EditorDebuggerSession.is_active].
*/
func (self Instance) GetSessions() gd.Array {
	return gd.Array(class(self).GetSessions())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorDebuggerPlugin

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorDebuggerPlugin"))
	return Instance{classdb.EditorDebuggerPlugin(object)}
}

/*
Override this method to be notified whenever a new [EditorDebuggerSession] is created (the session may be inactive during this stage).
*/
func (class) _setup_session(impl func(ptr unsafe.Pointer, session_id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var session_id = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, session_id)
	}
}

/*
Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passes to the [method _capture] method.
*/
func (class) _has_capture(impl func(ptr unsafe.Pointer, capture gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var capture = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, capture)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the message (which you can retrieve via [method get_session]).
*/
func (class) _capture(impl func(ptr unsafe.Pointer, message gd.String, data gd.Array, session_id gd.Int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var message = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var data = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var session_id = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, message, data, session_id)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
*/
func (class) _goto_script_line(impl func(ptr unsafe.Pointer, script objects.Script, line gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var script = objects.Script{pointers.New[classdb.Script]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(script[0])
		var line = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, line)
	}
}

/*
Override this method to be notified when all breakpoints are cleared in the editor.
*/
func (class) _breakpoints_cleared_in_tree(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to be notified when a breakpoint is set in the editor.
*/
func (class) _breakpoint_set_in_tree(impl func(ptr unsafe.Pointer, script objects.Script, line gd.Int, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var script = objects.Script{pointers.New[classdb.Script]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(script[0])
		var line = gd.UnsafeGet[gd.Int](p_args, 1)
		var enabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, line, enabled)
	}
}

/*
Returns the [EditorDebuggerSession] with the given [param id].
*/
//go:nosplit
func (self class) GetSession(id gd.Int) objects.EditorDebuggerSession {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerPlugin.Bind_get_session, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.EditorDebuggerSession{classdb.EditorDebuggerSession(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns an array of [EditorDebuggerSession] currently available to this debugger plugin.
[b]Note:[/b] Sessions in the array may be inactive, check their state via [method EditorDebuggerSession.is_active].
*/
//go:nosplit
func (self class) GetSessions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerPlugin.Bind_get_sessions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorDebuggerPlugin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorDebuggerPlugin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted         { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted      { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

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
		return gd.VirtualByName(self.AsRefCounted(), name)
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
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorDebuggerPlugin", func(ptr gd.Object) any { return classdb.EditorDebuggerPlugin(ptr) })
}
