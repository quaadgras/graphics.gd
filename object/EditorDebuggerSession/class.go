package EditorDebuggerSession

import "unsafe"
import "reflect"
import "runtime.link/mmm"
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
This class cannot be directly instantiated and must be retrieved via a [EditorDebuggerPlugin].
You can add tabs to the session UI via [method add_session_tab], send messages via [method send_message], and toggle [EngineProfiler]s via [method toggle_profiler].

*/
type Simple [1]classdb.EditorDebuggerSession
func (self Simple) SendMessage(message string, data gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SendMessage(gc.String(message), data)
}
func (self Simple) ToggleProfiler(profiler string, enable bool, data gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ToggleProfiler(gc.String(profiler), enable, data)
}
func (self Simple) IsBreaked() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBreaked())
}
func (self Simple) IsDebuggable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDebuggable())
}
func (self Simple) IsActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsActive())
}
func (self Simple) AddSessionTab(control [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddSessionTab(control)
}
func (self Simple) RemoveSessionTab(control [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveSessionTab(control)
}
func (self Simple) SetBreakpoint(path string, line int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBreakpoint(gc.String(path), gd.Int(line), enabled)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorDebuggerSession
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sends the given [param message] to the attached remote instance, optionally passing additionally [param data]. See [EngineDebugger] for how to retrieve those messages.
*/
//go:nosplit
func (self class) SendMessage(message gd.String, data gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_send_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Toggle the given [param profiler] on the attached remote instance, optionally passing additionally [param data]. See [EngineProfiler] for more details.
*/
//go:nosplit
func (self class) ToggleProfiler(profiler gd.String, enable bool, data gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(profiler))
	callframe.Arg(frame, enable)
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_toggle_profiler, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the attached remote instance is currently in the debug loop.
*/
//go:nosplit
func (self class) IsBreaked() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_is_breaked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the attached remote instance can be debugged.
*/
//go:nosplit
func (self class) IsDebuggable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_is_debuggable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the debug session is currently attached to a remote instance.
*/
//go:nosplit
func (self class) IsActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds the given [param control] to the debug session UI in the debugger bottom panel.
*/
//go:nosplit
func (self class) AddSessionTab(control object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(control[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_add_session_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the given [param control] from the debug session UI in the debugger bottom panel.
*/
//go:nosplit
func (self class) RemoveSessionTab(control object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(control[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_remove_session_tab, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables or disables a specific breakpoint based on [param enabled], updating the Editor Breakpoint Panel accordingly.
*/
//go:nosplit
func (self class) SetBreakpoint(path gd.String, line gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, line)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorDebuggerSession.Bind_set_breakpoint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEditorDebuggerSession() Expert { return self[0].AsEditorDebuggerSession() }


//go:nosplit
func (self Simple) AsEditorDebuggerSession() Simple { return self[0].AsEditorDebuggerSession() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorDebuggerSession", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
