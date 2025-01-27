// Package EditorDebuggerSession provides methods for working with EditorDebuggerSession object instances.
package EditorDebuggerSession

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
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"

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

/*
This class cannot be directly instantiated and must be retrieved via a [EditorDebuggerPlugin].
You can add tabs to the session UI via [method add_session_tab], send messages via [method send_message], and toggle [EngineProfiler]s via [method toggle_profiler].
*/
type Instance [1]gdclass.EditorDebuggerSession

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorDebuggerSession() Instance
}

/*
Sends the given [param message] to the attached remote instance, optionally passing additionally [param data]. See [EngineDebugger] for how to retrieve those messages.
*/
func (self Instance) SendMessage(message string) { //gd:EditorDebuggerSession.send_message
	class(self).SendMessage(String.New(message), Array.Nil)
}

/*
Toggle the given [param profiler] on the attached remote instance, optionally passing additionally [param data]. See [EngineProfiler] for more details.
*/
func (self Instance) ToggleProfiler(profiler string, enable bool) { //gd:EditorDebuggerSession.toggle_profiler
	class(self).ToggleProfiler(String.New(profiler), enable, Array.Nil)
}

/*
Returns [code]true[/code] if the attached remote instance is currently in the debug loop.
*/
func (self Instance) IsBreaked() bool { //gd:EditorDebuggerSession.is_breaked
	return bool(class(self).IsBreaked())
}

/*
Returns [code]true[/code] if the attached remote instance can be debugged.
*/
func (self Instance) IsDebuggable() bool { //gd:EditorDebuggerSession.is_debuggable
	return bool(class(self).IsDebuggable())
}

/*
Returns [code]true[/code] if the debug session is currently attached to a remote instance.
*/
func (self Instance) IsActive() bool { //gd:EditorDebuggerSession.is_active
	return bool(class(self).IsActive())
}

/*
Adds the given [param control] to the debug session UI in the debugger bottom panel.
*/
func (self Instance) AddSessionTab(control [1]gdclass.Control) { //gd:EditorDebuggerSession.add_session_tab
	class(self).AddSessionTab(control)
}

/*
Removes the given [param control] from the debug session UI in the debugger bottom panel.
*/
func (self Instance) RemoveSessionTab(control [1]gdclass.Control) { //gd:EditorDebuggerSession.remove_session_tab
	class(self).RemoveSessionTab(control)
}

/*
Enables or disables a specific breakpoint based on [param enabled], updating the Editor Breakpoint Panel accordingly.
*/
func (self Instance) SetBreakpoint(path string, line int, enabled bool) { //gd:EditorDebuggerSession.set_breakpoint
	class(self).SetBreakpoint(String.New(path), gd.Int(line), enabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorDebuggerSession

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorDebuggerSession"))
	casted := Instance{*(*gdclass.EditorDebuggerSession)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Sends the given [param message] to the attached remote instance, optionally passing additionally [param data]. See [EngineDebugger] for how to retrieve those messages.
*/
//go:nosplit
func (self class) SendMessage(message String.Readable, data Array.Any) { //gd:EditorDebuggerSession.send_message
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_send_message, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Toggle the given [param profiler] on the attached remote instance, optionally passing additionally [param data]. See [EngineProfiler] for more details.
*/
//go:nosplit
func (self class) ToggleProfiler(profiler String.Readable, enable bool, data Array.Any) { //gd:EditorDebuggerSession.toggle_profiler
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(profiler)))
	callframe.Arg(frame, enable)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_toggle_profiler, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the attached remote instance is currently in the debug loop.
*/
//go:nosplit
func (self class) IsBreaked() bool { //gd:EditorDebuggerSession.is_breaked
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_is_breaked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the attached remote instance can be debugged.
*/
//go:nosplit
func (self class) IsDebuggable() bool { //gd:EditorDebuggerSession.is_debuggable
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_is_debuggable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the debug session is currently attached to a remote instance.
*/
//go:nosplit
func (self class) IsActive() bool { //gd:EditorDebuggerSession.is_active
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds the given [param control] to the debug session UI in the debugger bottom panel.
*/
//go:nosplit
func (self class) AddSessionTab(control [1]gdclass.Control) { //gd:EditorDebuggerSession.add_session_tab
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(control[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_add_session_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the given [param control] from the debug session UI in the debugger bottom panel.
*/
//go:nosplit
func (self class) RemoveSessionTab(control [1]gdclass.Control) { //gd:EditorDebuggerSession.remove_session_tab
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_remove_session_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Enables or disables a specific breakpoint based on [param enabled], updating the Editor Breakpoint Panel accordingly.
*/
//go:nosplit
func (self class) SetBreakpoint(path String.Readable, line gd.Int, enabled bool) { //gd:EditorDebuggerSession.set_breakpoint
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, line)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_set_breakpoint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnStarted(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnStopped(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("stopped"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBreaked(cb func(can_debug bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("breaked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnContinued(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("continued"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorDebuggerSession() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorDebuggerSession() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorDebuggerSession", func(ptr gd.Object) any {
		return [1]gdclass.EditorDebuggerSession{*(*gdclass.EditorDebuggerSession)(unsafe.Pointer(&ptr))}
	})
}
