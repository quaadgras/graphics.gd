package EditorDebuggerPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

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
		SetupSession(session_id gd.Int) 
		//Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passes to the [method _capture] method.
		HasCapture(capture gd.String) bool
		//Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the message (which you can retrieve via [method get_session]).
		Capture(message gd.String, data gd.Array, session_id gd.Int) bool
		//Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
		GotoScriptLine(script object.Script, line gd.Int) 
		//Override this method to be notified when all breakpoints are cleared in the editor.
		BreakpointsClearedInTree() 
		//Override this method to be notified when a breakpoint is set in the editor.
		BreakpointSetInTree(script object.Script, line gd.Int, enabled bool) 
	}

*/
type Simple [1]classdb.EditorDebuggerPlugin
func (Simple) _setup_session(impl func(ptr unsafe.Pointer, session_id int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var session_id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(session_id))
		gc.End()
	}
}
func (Simple) _has_capture(impl func(ptr unsafe.Pointer, capture string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var capture = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, capture.String())
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _capture(impl func(ptr unsafe.Pointer, message string, data gd.Array, session_id int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var message = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var data = mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var session_id = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, message.String(), data, int(session_id))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _goto_script_line(impl func(ptr unsafe.Pointer, script [1]classdb.Script, line int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var script [1]classdb.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var line = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, script, int(line))
		gc.End()
	}
}
func (Simple) _breakpoints_cleared_in_tree(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _breakpoint_set_in_tree(impl func(ptr unsafe.Pointer, script [1]classdb.Script, line int, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var script [1]classdb.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var line = gd.UnsafeGet[gd.Int](p_args,1)
		var enabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, script, int(line), enabled)
		gc.End()
	}
}
func (self Simple) GetSession(id int) [1]classdb.EditorDebuggerSession {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.EditorDebuggerSession(Expert(self).GetSession(gc, gd.Int(id)))
}
func (self Simple) GetSessions() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Array(Expert(self).GetSessions(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorDebuggerPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Override this method to be notified whenever a new [EditorDebuggerSession] is created (the session may be inactive during this stage).
*/
func (class) _setup_session(impl func(ptr unsafe.Pointer, session_id gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var session_id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, session_id)
		ctx.End()
	}
}

/*
Override this method to enable receiving messages from the debugger. If [param capture] is "my_message" then messages starting with "my_message:" will be passes to the [method _capture] method.
*/
func (class) _has_capture(impl func(ptr unsafe.Pointer, capture gd.String) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var capture = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, capture)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to process incoming messages. The [param session_id] is the ID of the [EditorDebuggerSession] that received the message (which you can retrieve via [method get_session]).
*/
func (class) _capture(impl func(ptr unsafe.Pointer, message gd.String, data gd.Array, session_id gd.Int) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var message = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var data = mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var session_id = gd.UnsafeGet[gd.Int](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, message, data, session_id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to be notified when a breakpoint line has been clicked in the debugger breakpoint panel.
*/
func (class) _goto_script_line(impl func(ptr unsafe.Pointer, script object.Script, line gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var script object.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var line = gd.UnsafeGet[gd.Int](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, script, line)
		ctx.End()
	}
}

/*
Override this method to be notified when all breakpoints are cleared in the editor.
*/
func (class) _breakpoints_cleared_in_tree(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Override this method to be notified when a breakpoint is set in the editor.
*/
func (class) _breakpoint_set_in_tree(impl func(ptr unsafe.Pointer, script object.Script, line gd.Int, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var script object.Script
		script[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var line = gd.UnsafeGet[gd.Int](p_args,1)
		var enabled = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, script, line, enabled)
		ctx.End()
	}
}

/*
Returns the [EditorDebuggerSession] with the given [param id].
*/
//go:nosplit
func (self class) GetSession(ctx gd.Lifetime, id gd.Int) object.EditorDebuggerSession {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerPlugin.Bind_get_session, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.EditorDebuggerSession
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an array of [EditorDebuggerSession] currently available to this debugger plugin.
[b]Note:[/b] Sessions in the array may be inactive, check their state via [method EditorDebuggerSession.is_active].
*/
//go:nosplit
func (self class) GetSessions(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerPlugin.Bind_get_sessions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorDebuggerPlugin() Expert { return self[0].AsEditorDebuggerPlugin() }


//go:nosplit
func (self Simple) AsEditorDebuggerPlugin() Simple { return self[0].AsEditorDebuggerPlugin() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_session": return reflect.ValueOf(self._setup_session);
	case "_has_capture": return reflect.ValueOf(self._has_capture);
	case "_capture": return reflect.ValueOf(self._capture);
	case "_goto_script_line": return reflect.ValueOf(self._goto_script_line);
	case "_breakpoints_cleared_in_tree": return reflect.ValueOf(self._breakpoints_cleared_in_tree);
	case "_breakpoint_set_in_tree": return reflect.ValueOf(self._breakpoint_set_in_tree);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_setup_session": return reflect.ValueOf(self._setup_session);
	case "_has_capture": return reflect.ValueOf(self._has_capture);
	case "_capture": return reflect.ValueOf(self._capture);
	case "_goto_script_line": return reflect.ValueOf(self._goto_script_line);
	case "_breakpoints_cleared_in_tree": return reflect.ValueOf(self._breakpoints_cleared_in_tree);
	case "_breakpoint_set_in_tree": return reflect.ValueOf(self._breakpoint_set_in_tree);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorDebuggerPlugin", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
