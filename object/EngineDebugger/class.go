package EngineDebugger

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
[EngineDebugger] handles the communication between the editor and the running game. It is active in the running game. Messages can be sent/received through it. It also manages the profilers.

*/
type Simple [1]classdb.EngineDebugger
func (self Simple) IsActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsActive())
}
func (self Simple) RegisterProfiler(name string, profiler [1]classdb.EngineProfiler) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterProfiler(gc.StringName(name), profiler)
}
func (self Simple) UnregisterProfiler(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnregisterProfiler(gc.StringName(name))
}
func (self Simple) IsProfiling(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsProfiling(gc.StringName(name)))
}
func (self Simple) HasProfiler(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasProfiler(gc.StringName(name)))
}
func (self Simple) ProfilerAddFrameData(name string, data gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ProfilerAddFrameData(gc.StringName(name), data)
}
func (self Simple) ProfilerEnable(name string, enable bool, arguments gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ProfilerEnable(gc.StringName(name), enable, arguments)
}
func (self Simple) RegisterMessageCapture(name string, callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterMessageCapture(gc.StringName(name), callable)
}
func (self Simple) UnregisterMessageCapture(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnregisterMessageCapture(gc.StringName(name))
}
func (self Simple) HasCapture(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasCapture(gc.StringName(name)))
}
func (self Simple) LinePoll() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).LinePoll()
}
func (self Simple) SendMessage(message string, data gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SendMessage(gc.String(message), data)
}
func (self Simple) Debug(can_continue bool, is_error_breakpoint bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Debug(can_continue, is_error_breakpoint)
}
func (self Simple) ScriptDebug(language [1]classdb.ScriptLanguage, can_continue bool, is_error_breakpoint bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ScriptDebug(language, can_continue, is_error_breakpoint)
}
func (self Simple) SetLinesLeft(lines int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLinesLeft(gd.Int(lines))
}
func (self Simple) GetLinesLeft() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLinesLeft()))
}
func (self Simple) SetDepth(depth int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDepth(gd.Int(depth))
}
func (self Simple) GetDepth() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetDepth()))
}
func (self Simple) IsBreakpoint(line int, source string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBreakpoint(gd.Int(line), gc.StringName(source)))
}
func (self Simple) IsSkippingBreakpoints() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSkippingBreakpoints())
}
func (self Simple) InsertBreakpoint(line int, source string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).InsertBreakpoint(gd.Int(line), gc.StringName(source))
}
func (self Simple) RemoveBreakpoint(line int, source string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveBreakpoint(gd.Int(line), gc.StringName(source))
}
func (self Simple) ClearBreakpoints() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearBreakpoints()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EngineDebugger
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns [code]true[/code] if the debugger is active otherwise [code]false[/code].
*/
//go:nosplit
func (self class) IsActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Registers a profiler with the given [param name]. See [EngineProfiler] for more information.
*/
//go:nosplit
func (self class) RegisterProfiler(name gd.StringName, profiler object.EngineProfiler)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(profiler[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_register_profiler, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters a profiler with given [param name].
*/
//go:nosplit
func (self class) UnregisterProfiler(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_unregister_profiler, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a profiler with the given name is present and active otherwise [code]false[/code].
*/
//go:nosplit
func (self class) IsProfiling(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_is_profiling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a profiler with the given name is present otherwise [code]false[/code].
*/
//go:nosplit
func (self class) HasProfiler(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_has_profiler, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Calls the [code]add[/code] callable of the profiler with given [param name] and [param data].
*/
//go:nosplit
func (self class) ProfilerAddFrameData(name gd.StringName, data gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_profiler_add_frame_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Calls the [code]toggle[/code] callable of the profiler with given [param name] and [param arguments]. Enables/Disables the same profiler depending on [param enable] argument.
*/
//go:nosplit
func (self class) ProfilerEnable(name gd.StringName, enable bool, arguments gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, enable)
	callframe.Arg(frame, mmm.Get(arguments))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_profiler_enable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers a message capture with given [param name]. If [param name] is "my_message" then messages starting with "my_message:" will be called with the given callable.
Callable must accept a message string and a data array as argument. If the message and data are valid then callable must return [code]true[/code] otherwise [code]false[/code].
*/
//go:nosplit
func (self class) RegisterMessageCapture(name gd.StringName, callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_register_message_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters the message capture with given [param name].
*/
//go:nosplit
func (self class) UnregisterMessageCapture(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_unregister_message_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a capture with the given name is present otherwise [code]false[/code].
*/
//go:nosplit
func (self class) HasCapture(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_has_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces a processing loop of debugger events. The purpose of this method is just processing events every now and then when the script might get too busy, so that bugs like infinite loops can be caught
*/
//go:nosplit
func (self class) LinePoll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_line_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sends a message with given [param message] and [param data] array.
*/
//go:nosplit
func (self class) SendMessage(message gd.String, data gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(message))
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_send_message, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Starts a debug break in script execution, optionally specifying whether the program can continue based on [param can_continue] and whether the break was due to a breakpoint.
*/
//go:nosplit
func (self class) Debug(can_continue bool, is_error_breakpoint bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, can_continue)
	callframe.Arg(frame, is_error_breakpoint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_debug, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Starts a debug break in script execution, optionally specifying whether the program can continue based on [param can_continue] and whether the break was due to a breakpoint.
*/
//go:nosplit
func (self class) ScriptDebug(language object.ScriptLanguage, can_continue bool, is_error_breakpoint bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(language[0].AsPointer())[0])
	callframe.Arg(frame, can_continue)
	callframe.Arg(frame, is_error_breakpoint)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_script_debug, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the current debugging lines that remain.
*/
//go:nosplit
func (self class) SetLinesLeft(lines gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, lines)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_set_lines_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of lines that remain.
*/
//go:nosplit
func (self class) GetLinesLeft() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_get_lines_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the current debugging depth.
*/
//go:nosplit
func (self class) SetDepth(depth gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current debug depth.
*/
//go:nosplit
func (self class) GetDepth() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the given [param source] and [param line] represent an existing breakpoint.
*/
//go:nosplit
func (self class) IsBreakpoint(line gd.Int, source gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, mmm.Get(source))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_is_breakpoint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the debugger is skipping breakpoints otherwise [code]false[/code].
*/
//go:nosplit
func (self class) IsSkippingBreakpoints() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_is_skipping_breakpoints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Inserts a new breakpoint with the given [param source] and [param line].
*/
//go:nosplit
func (self class) InsertBreakpoint(line gd.Int, source gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, mmm.Get(source))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_insert_breakpoint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a breakpoint with the given [param source] and [param line].
*/
//go:nosplit
func (self class) RemoveBreakpoint(line gd.Int, source gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, mmm.Get(source))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_remove_breakpoint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all breakpoints.
*/
//go:nosplit
func (self class) ClearBreakpoints()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EngineDebugger.Bind_clear_breakpoints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsEngineDebugger() Expert { return self[0].AsEngineDebugger() }


//go:nosplit
func (self Simple) AsEngineDebugger() Simple { return self[0].AsEngineDebugger() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EngineDebugger", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
