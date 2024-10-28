package EditorDebuggerSession

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class cannot be directly instantiated and must be retrieved via a [EditorDebuggerPlugin].
You can add tabs to the session UI via [method add_session_tab], send messages via [method send_message], and toggle [EngineProfiler]s via [method toggle_profiler].

*/
type Go [1]classdb.EditorDebuggerSession

/*
Sends the given [param message] to the attached remote instance, optionally passing additionally [param data]. See [EngineDebugger] for how to retrieve those messages.
*/
func (self Go) SendMessage(message string) {
	class(self).SendMessage(gd.NewString(message), ([1]gd.Array{}[0]))
}

/*
Toggle the given [param profiler] on the attached remote instance, optionally passing additionally [param data]. See [EngineProfiler] for more details.
*/
func (self Go) ToggleProfiler(profiler string, enable bool) {
	class(self).ToggleProfiler(gd.NewString(profiler), enable, ([1]gd.Array{}[0]))
}

/*
Returns [code]true[/code] if the attached remote instance is currently in the debug loop.
*/
func (self Go) IsBreaked() bool {
	return bool(class(self).IsBreaked())
}

/*
Returns [code]true[/code] if the attached remote instance can be debugged.
*/
func (self Go) IsDebuggable() bool {
	return bool(class(self).IsDebuggable())
}

/*
Returns [code]true[/code] if the debug session is currently attached to a remote instance.
*/
func (self Go) IsActive() bool {
	return bool(class(self).IsActive())
}

/*
Adds the given [param control] to the debug session UI in the debugger bottom panel.
*/
func (self Go) AddSessionTab(control gdclass.Control) {
	class(self).AddSessionTab(control)
}

/*
Removes the given [param control] from the debug session UI in the debugger bottom panel.
*/
func (self Go) RemoveSessionTab(control gdclass.Control) {
	class(self).RemoveSessionTab(control)
}

/*
Enables or disables a specific breakpoint based on [param enabled], updating the Editor Breakpoint Panel accordingly.
*/
func (self Go) SetBreakpoint(path string, line int, enabled bool) {
	class(self).SetBreakpoint(gd.NewString(path), gd.Int(line), enabled)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorDebuggerSession
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorDebuggerSession"))
	return Go{classdb.EditorDebuggerSession(object)}
}

/*
Sends the given [param message] to the attached remote instance, optionally passing additionally [param data]. See [EngineDebugger] for how to retrieve those messages.
*/
//go:nosplit
func (self class) SendMessage(message gd.String, data gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(message))
	callframe.Arg(frame, discreet.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_send_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Toggle the given [param profiler] on the attached remote instance, optionally passing additionally [param data]. See [EngineProfiler] for more details.
*/
//go:nosplit
func (self class) ToggleProfiler(profiler gd.String, enable bool, data gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(profiler))
	callframe.Arg(frame, enable)
	callframe.Arg(frame, discreet.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_toggle_profiler, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the attached remote instance is currently in the debug loop.
*/
//go:nosplit
func (self class) IsBreaked() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_is_breaked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the attached remote instance can be debugged.
*/
//go:nosplit
func (self class) IsDebuggable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_is_debuggable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the debug session is currently attached to a remote instance.
*/
//go:nosplit
func (self class) IsActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the given [param control] to the debug session UI in the debugger bottom panel.
*/
//go:nosplit
func (self class) AddSessionTab(control gdclass.Control)  {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(control[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_add_session_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given [param control] from the debug session UI in the debugger bottom panel.
*/
//go:nosplit
func (self class) RemoveSessionTab(control gdclass.Control)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(control[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_remove_session_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables or disables a specific breakpoint based on [param enabled], updating the Editor Breakpoint Panel accordingly.
*/
//go:nosplit
func (self class) SetBreakpoint(path gd.String, line gd.Int, enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	callframe.Arg(frame, line)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorDebuggerSession.Bind_set_breakpoint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnStarted(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("started"), gd.NewCallable(cb), 0)
}


func (self Go) OnStopped(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("stopped"), gd.NewCallable(cb), 0)
}


func (self Go) OnBreaked(cb func(can_debug bool)) {
	self[0].AsObject().Connect(gd.NewStringName("breaked"), gd.NewCallable(cb), 0)
}


func (self Go) OnContinued(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("continued"), gd.NewCallable(cb), 0)
}


func (self class) AsEditorDebuggerSession() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorDebuggerSession() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorDebuggerSession", func(ptr gd.Object) any { return classdb.EditorDebuggerSession(ptr) })}
